package main

import (
	"lms-backend/internal/app"
)

// @title Library Mangement System API
func main() {
	// Setup and run the app
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
