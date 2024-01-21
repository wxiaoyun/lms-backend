.PHONY: docs

clean:
	go mod tidy
	go mod vendor

test:
	go test -v ./...

createdb:
	go run cmd/createdb/main.go

dropdb:
	go run cmd/dropdb/main.go

migratedb:
	go run cmd/migratedb/main.go -dir=up

rollbackdb:
	go run cmd/migratedb/main.go -dir=down -step=$(step) 

seeddb:
	go run cmd/seeddb/main.go

flushdb:
	go run cmd/flushdb/main.go

setupDB: createDB migrateDB seedDB

resetDB: dropDB setupDB

count:
	git ls-files '*.go' | grep -v '^docs/' | xargs wc -l

dockerbuild:
	docker compose --env-file .env.development build

dockerup:
	docker compose --env-file .env.development up

dockerdown:
	docker compose --env-file .env.development down

dockershell:
	docker exec -it lms-backend-server-1 sh