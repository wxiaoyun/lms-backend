package storage

import (
	"lms-backend/pkg/error/externalerrors"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Storage manages file operations within a specific base directory.
type Storage struct {
	// BaseDirectoryElems is a slice of strings representing the elements of the base directory.
	BaseDirectoryElems []string
}

// New creates a new Storage instance from the provided slice of directory elements.
func New(baseDirElems ...string) *Storage {
	return &Storage{BaseDirectoryElems: baseDirElems}
}

// constructFilePath constructs a safe file path within the base directory from the given path elements.
func (s *Storage) ConstructFilePath(pathElems ...string) (string, error) {
	// Sanitize individual path elements to prevent directory traversal
	for i, elem := range pathElems {
		// Clean the path element to remove any relative elements or suspicious characters
		cleanElem := filepath.Clean(elem)
		// Prevent directory traversal by ensuring that the cleaned element doesn't contain
		// path traversal segments like ".."
		if cleanElem == ".." || strings.Contains(cleanElem, ".."+string(os.PathSeparator)) {
			return "", externalerrors.BadRequest("invalid path element: " + elem)
		}
		pathElems[i] = cleanElem
	}

	// Construct the full path
	allElems := s.BaseDirectoryElems
	allElems = append(allElems, pathElems...)
	cleanPath := filepath.Join(allElems...)

	// Convert to absolute path
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return "", errors.Wrap(err, "resolving absolute path failed")
	}

	// Ensure the path is within the base directory
	basePath := filepath.Join(s.BaseDirectoryElems...)
	baseAbsPath, err := filepath.Abs(basePath)
	if err != nil {
		return "", errors.Wrap(err, "resolving absolute path for base directory failed")
	}

	if !strings.HasPrefix(absPath, baseAbsPath) {
		return "", externalerrors.BadRequest("path traversal detected")
	}

	return absPath, nil
}

func (s *Storage) ValidateFilePath(path string) error {
	// Clean the path element to remove any relative elements or suspicious characters
	cleanElem := filepath.Clean(path)
	// Prevent directory traversal by ensuring that the cleaned element doesn't contain
	// path traversal segments like ".."
	if cleanElem == ".." || strings.Contains(cleanElem, ".."+string(os.PathSeparator)) {
		return externalerrors.BadRequest("invalid path: " + path)
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return errors.Wrap(err, "resolving absolute path failed")
	}

	// Ensure the path is within the base directory
	basePath := filepath.Join(s.BaseDirectoryElems...)
	baseAbsPath, err := filepath.Abs(basePath)
	if err != nil {
		return errors.Wrap(err, "resolving absolute path for base directory failed")
	}

	if !strings.HasPrefix(absPath, baseAbsPath) {
		return externalerrors.BadRequest("path traversal detected")
	}

	return nil
}
