package main

import "auth-practice/internal/app"

func main() {
	// Setup and run the app
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
