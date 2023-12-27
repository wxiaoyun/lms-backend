.PHONY: docs

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

runRedis:
	docker run --hostname=2ea46ec4c044 --env=PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin --env=REDISBLOOM_ARGS= --env=REDISEARCH_ARGS= --env=REDISJSON_ARGS= --env=REDISTIMESERIES_ARGS= --env=REDIS_ARGS= -p 6379:6379 -p 8001:8001 --restart=no --label='org.opencontainers.image.ref.name=ubuntu' --label='org.opencontainers.image.version=22.04' --runtime=runc -d redis/redis-stack:latest

count:
	git ls-files '*.go' | grep -v '^docs/' | xargs wc -l