run:
	go run cmd/calljournal/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o calljournal cmd/calljournal/main.go

runcdrserver:
	go run cmd/cdrserver/main.go

runcdrclient:
	go run cmd/cdrclient/main.go

runcdrstore:
	go run cmd/cdrstore/main.go

migrateup:
	goose -dir migrations postgres "user=${CJ_DB_USER} password=${CJ_DB_PASSWORD} dbname=${CJ_DB_NAME} sslmode=disable" up

migratedown:
	goose -dir migrations postgres "user=${CJ_DB_USER} password=${CJ_DB_PASSWORD} dbname=${CJ_DB_NAME} sslmode=disable" down

lint:
	golangci-lint run ./...

test:
	go test -v -race ./... -count=1

integration-tests:
	docker-compose -f ./docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from integration_tests && \
	docker-compose -f ./docker-compose.test.yml down

generate:
	go generate proto/gen.go
	go generate bindata/staticfs/gen.go

template:
	go generate bindata/tmpl/gen.go

clean:
	rm -rf internal/pb/*

.PHONY: run template runcdrserver runcdrclient runcdrstore migratedown migrateup lint generate clean intergration-tests