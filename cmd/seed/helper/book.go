package shelper

import (
	"fmt"
	"lms-backend/internal/model"
	"lms-backend/util/random"
	"strings"
	"time"

	"gorm.io/gorm"
)

var (
	languages = []string{
		"en",
		"km",
	}
)

func randomLanguage() string {
	return languages[random.RandInt(0, len(languages))]
}

func escapeSQL(value string) string {
	// Basic escaping for single quotes, replace with more comprehensive escaping if needed
	return strings.ReplaceAll(value, "'", "''")
}

func GenerateBulkInsertBooksSQL(num int64) string {
	var valueStrings []string

	for i := 0; i < int(num); i++ {
		// Generate book data directly in the format required by SQL
		title := escapeSQL(random.RandWords(random.RandInt(4, 11)))
		author := escapeSQL(random.RandWords(random.RandInt(2, 5)))
		isbn := escapeSQL(GenerateISBN13())
		publisher := escapeSQL(random.RandWords(random.RandInt(4, 7)))
		publicationDate := random.RandomDate(time.Now().AddDate(-10, 0, 0), time.Now()).Format("2006-01-02")
		genre := escapeSQL(random.RandWords(random.RandInt(1, 3)))
		language := escapeSQL(randomLanguage())
		createdAt := time.Now().Format("2006-01-02 15:04:05")

		valueString := fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", title, author, isbn, publisher, publicationDate, genre, language, createdAt)
		valueStrings = append(valueStrings, valueString)
	}

	stmt := fmt.Sprintf("INSERT INTO books (title, author, isbn, publisher, publication_date, genre, language, created_at) VALUES %s", strings.Join(valueStrings, ","))
	return stmt
}

func GenerateBulkInsertBookCopiesSQL(count, max int64) string {
	var valueStrings []string

	for i := 0; i < int(count); i++ {
		numOfCopies := random.RandInt(1, int(max+1))
		for j := 0; j < numOfCopies; j++ {
			status := model.BookStatusAvailable
			createdAt := time.Now().Format("2006-01-02 15:04:05")

			valueString := fmt.Sprintf("(%d, '%s', '%s')", i+1, status, createdAt)
			valueStrings = append(valueStrings, valueString)
		}
	}

	return fmt.Sprintf("INSERT INTO book_copies (book_id, status, created_at) VALUES %s", strings.Join(valueStrings, ","))
}

func SeedBookAndCopies(db *gorm.DB, num int64) error {
	var count int64

	result := db.Model(&model.Book{}).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count >= num {
		return nil
	}

	stmt := GenerateBulkInsertBooksSQL(num - count)
	err := db.Exec(stmt).Error
	if err != nil {
		return err
	}

	stmt = GenerateBulkInsertBookCopiesSQL(num, 5)
	err = db.Exec(stmt).Error
	if err != nil {
		return err
	}

	return nil
}
