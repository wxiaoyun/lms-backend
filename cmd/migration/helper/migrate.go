package migration

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"technical-test/internal/config"
	"technical-test/internal/database"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Runs the .sql migration file from <filename> and  "-- +migrate <section>" section
//
// Example:
//
// -- migrate up
//
//	CREATE TABLE users (
//		id SERIAL PRIMARY KEY,
//		name VARCHAR(255)
//	);
//
// -- migrate down
//
//	DROP TABLE users;
func RunMigration(filename, section string) error {
	//nolint:gosec // ignore error
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	var inSection bool
	var sqlCommands []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, fmt.Sprintf("-- +migrate %s", section)) {
			inSection = !inSection
			continue
		} else if strings.Contains(line, "-- +migrate") {
			// Any other "-- migrate" line will stop the current section
			inSection = false
		}

		if inSection {
			sqlCommands = append(sqlCommands, line)
		}
	}

	if len(sqlCommands) == 0 {
		return fmt.Errorf("no SQL commands found in section %s", section)
	}

	// Join all lines to form the full SQL content
	fullSQL := strings.Join(sqlCommands, " ")

	// Split the content by semicolons to get individual SQL commands
	individualCommands := strings.Split(fullSQL, ";")

	cf, err := config.LoadEnvAndGetConfig()
	if err != nil {
		return err
	}

	dsn, err := database.GormDSNBuilder(cf)
	if err != nil {
		return err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
			fmt.Errorf("Unsuccessful migration: %s\n", filename)
		} else {
			tx.Commit()
			//nolint:revive // ignore lint
			fmt.Printf("Migrated successfully file: %s\n", filename)
		}
	}()

	// Execute SQL commands
	for _, cmd := range individualCommands {
		trimmedCmd := strings.TrimSpace(cmd)
		if trimmedCmd != "" {
			if err = tx.Exec(trimmedCmd).Error; err != nil {
				fmt.Errorf("error executing SQL command: %s", err)
				panic(err) // panic to rollback transaction
			}
		}
	}

	return nil
}
