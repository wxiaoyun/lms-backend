package migration

import (
	"os"
	"path/filepath"
	"strings"
)

// Retrieve all .sql files from a specified directory:
func ListSQLFiles(directory string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".sql") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
