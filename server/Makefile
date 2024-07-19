SHELL := env PATH=\"$(PATH)\" bash

.PHONY: migrate-up
migrate-up:
	migrate -path ./migrations -database "mysql://tweet:passwd@tcp(localhost:13306)/tweet" up
	migrate -path ./migrations -database "mysql://tweet:passwd@tcp(localhost:23306)/tweet" up

.PHONY: migrate-down
migrate-down:
	migrate -path ./migrations -database "mysql://tweet:passwd@tcp(localhost:13306)/tweet" down
	migrate -path ./migrations -database "mysql://tweet:passwd@tcp(localhost:23306)/tweet" down

.PHONY: dev
dev:
	ENV=local air main.go

.PHONY: test
test:
	ENV=test go test -v ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: seed
seed:
	go run cmd/seeder/main.go

.PHONY: build
build:
	docker buildx build -t tweet-server -f ./Dockerfile .
