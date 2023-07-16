clean:
	go mod tidy
	go mod vendor

run:
	go run main.go

migrateDB:
	go run cmd/migration/main.go