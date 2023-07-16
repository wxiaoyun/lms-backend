package main

import "auth-practice/internal/database"

func main() {
	database.OpenDataBase()
	database.AutoMigration()
}
