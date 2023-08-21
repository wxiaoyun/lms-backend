clean:
	go mod tidy
	go mod vendor

docs:
	swag init
	
run:
	go run main.go

test:
	go test -v ./...

createDB:
	go run cmd/createdb/main.go

dropDB:
	go run cmd/dropdb/main.go

migrateDB:
	go run cmd/migration/main.go -dir=up

rollbackDB:
	go run cmd/migration/main.go -dir=down -step=$(step) 

seedDB:
	go run cmd/seed/main.go

setupDB: createDB migrateDB seedDB

resetDB: dropDB setupDB

countlines:
	git ls-files | xargs wc -l