package bookparams

import (
	"lms-backend/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BaseParams struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	Publisher       string `json:"publisher"`
	PublicationDate string `json:"publication_date"`
	Genre           string `json:"genre"`
	Language        string `json:"language"`
}

func (p *BaseParams) Validate() error {
	if p.Title == "" {
		return fiber.NewError(fiber.StatusBadRequest, "title is required")
	}

	if p.Author == "" {
		return fiber.NewError(fiber.StatusBadRequest, "author is required")
	}

	if p.ISBN == "" {
		return fiber.NewError(fiber.StatusBadRequest, "isbn is required")
	}

	if p.Publisher == "" {
		return fiber.NewError(fiber.StatusBadRequest, "publisher is required")
	}

	if p.PublicationDate == "" {
		return fiber.NewError(fiber.StatusBadRequest, "publication_date is required")
	}

	if _, err := time.Parse(time.RFC3339, p.PublicationDate); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "publication_date does not match RFC3339 format")
	}

	if p.Genre == "" {
		return fiber.NewError(fiber.StatusBadRequest, "genre is required")
	}

	if p.Language == "" {
		return fiber.NewError(fiber.StatusBadRequest, "language is required")
	}

	return nil
}

func (p *BaseParams) ToModel() *model.Book {
	//nolint // err is checked in Validate()
	publicationDate, _ := time.Parse(time.RFC3339, p.PublicationDate)
	return &model.Book{
		Title:           p.Title,
		Author:          p.Author,
		ISBN:            p.ISBN,
		Publisher:       p.Publisher,
		PublicationDate: publicationDate,
		Genre:           p.Genre,
		Language:        p.Language,
	}
}
