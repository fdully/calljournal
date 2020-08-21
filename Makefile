run:
	go run cmd/calljournal/main.go

runuploader:
	go run cmd/calluploader/main.go

migrateup:
	goose -dir migrations postgres "user=${CJ_DB_USER} password=${CJ_DB_PASSWORD} dbname=${CJ_DB_NAME} sslmode=disable" up

migratedown:
	goose -dir migrations postgres "user=${CJ_DB_USER} password=${CJ_DB_PASSWORD} dbname=${CJ_DB_NAME} sslmode=disable" down

lint:
	golangci-lint run ./...

test:
	go test -v -race ./... -count=1

integration-tests:
	docker-compose -f ./docker-compose-tests.yaml up --build --abort-on-container-exit --exit-code-from integration_tests && \
	docker-compose -f ./docker-compose-tests.yaml down

generate:
	go generate proto/gen.go

clean:
	rm -rf internal/pb/*

.PHONY: run runuploader migratedown migrateup lint generate clean intergration-tests