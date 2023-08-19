package cdbhelper

import (
	"lms-backend/internal/config"
	"lms-backend/internal/database"
	"lms-backend/internal/policy/abilities"
	"lms-backend/internal/policy/roles"
)

func SeedDB(cf *config.Config) error {
	lgr.Println("Seeding database...")

	lgr.Println("GORM connecting to database...")
	err := database.OpenDataBase(cf)
	if err != nil {
		return err
	}
	lgr.Println("GORM connected to database.")

	db := database.GetDB()

	lgr.Println("Seeding Roles and Abilities...")
	db.Create(abilities.GetAllAbilities())
	db.Create(roles.GetAllRoles())

	return nil
}
