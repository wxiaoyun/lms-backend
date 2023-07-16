clean:
	go mod tidy
	go mod vendor

# Run the application with live reload
run:
	go run main.go

migrateDB:
	go run cmd/migration/main.go