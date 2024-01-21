.PHONY: docs

clean:
	go mod tidy
	go mod vendor

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

count:
	git ls-files '*.go' | grep -v '^docs/' | xargs wc -l

dockerbuild:
	docker compose --env-file .env.development build

dockerup:
	docker compose --env-file .env.development up

dockerdown:
	docker compose --env-file .env.development down

dockerterminal:
	docker exec -it lms-backend-server-1 sh