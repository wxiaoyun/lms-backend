clean:
	go mod tidy
	go mod vendor

swagger:
	swag init
	
# Run the application with live reload
run:
	go run main.go

createDB:
	go run cmd/createdb/main.go

dropDB:
	go run cmd/dropdb/main.go

migrateDB:
	go run cmd/migration/main.go

seedDB:
	go run cmd/seed/main.go

setupDB: createDB migrateDB seedDB