package model

import (
	"lms-backend/internal/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestWorksheetModel(t *testing.T) {
	db, err := testutil.ConnectToDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	db = db.Begin()

	// Positive test for Create
	t.Run("Create Valid Worksheet", func(t *testing.T) {
		worksheet := &Worksheet{Title: "Test worksheet", UserID: 1, Cost: 10.0, Price: 15.0, Description: "Test description"}
		err := worksheet.Create(db)
		assert.NoError(t, err)
	})

	// Negative test for Create
	t.Run("Create Invalid Worksheet", func(t *testing.T) {
		worksheet := &Worksheet{Title: "", UserID: 0, Cost: -5.0, Price: -10.0, Description: ""}
		err := worksheet.Create(db)
		assert.Error(t, err)
	})

	// Positive test for Update
	t.Run("Update Valid Worksheet", func(t *testing.T) {
		worksheet := &Worksheet{Model: gorm.Model{ID: 1}, Title: "Updated worksheet", UserID: 1, Cost: 15.0, Price: 20.0, Description: "Updated description"}
		err := worksheet.Update(db)
		assert.NoError(t, err)
	})

	// Negative test for Update
	t.Run("Update Invalid Worksheet", func(t *testing.T) {
		worksheet := &Worksheet{Model: gorm.Model{ID: 1}, Title: "", UserID: 0, Cost: -5.0, Price: -10.0, Description: ""}
		err := worksheet.Update(db)
		assert.Error(t, err)
	})

	db.Rollback()
}
