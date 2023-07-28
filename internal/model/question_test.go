package model

import (
	"technical-test/internal/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestQuestionModel(t *testing.T) {
	db, err := testutil.ConnectToDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	db = db.Begin()

	// Positive test for Create
	t.Run("Create Valid Question", func(t *testing.T) {
		question := &Question{Description: "Test question", Answer: "Test answer", Cost: 10.0, WorksheetID: 1}
		err := question.Create(db)
		assert.NoError(t, err)
	})

	// Negative test for Create
	t.Run("Create Invalid Question", func(t *testing.T) {
		question := &Question{Description: "", Answer: "", Cost: -5.0, WorksheetID: 0}
		err := question.Create(db)
		assert.Error(t, err)
	})

	// Positive test for Update
	t.Run("Update Valid Question", func(t *testing.T) {
		question := &Question{Model: gorm.Model{ID: 1}, Description: "Updated question", Answer: "Updated answer", Cost: 15.0, WorksheetID: 1}
		err := question.Update(db)
		assert.NoError(t, err)
	})

	// Negative test for Update
	t.Run("Update Invalid Question", func(t *testing.T) {
		question := &Question{Model: gorm.Model{ID: 1}, Description: "", Answer: "", Cost: -5.0, WorksheetID: 0}
		err := question.Update(db)
		assert.Error(t, err)
	})

	db.Rollback()
}
