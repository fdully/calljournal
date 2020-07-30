run:
	go run cmd/calljournal/main.go

migrateup:
	goose -dir migrations postgres "user=${CJ_DB_LOGIN} password=${CJ_DB_PASSWORD} dbname=${CJ_DB_NAME} sslmode=disable" up

migratedown:
	goose -dir migrations postgres "user=${CJ_DB_LOGIN} password=${CJ_DB_PASSWORD} dbname=${CJ_DB_NAME} sslmode=disable" down

lint:
	golangci-lint run ./...

test:
	go test -v -race ./... -count=1

generate:
	go generate proto/gen.go

clean:
	rm -rf internal/pb/*

.PHONY: