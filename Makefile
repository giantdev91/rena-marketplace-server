MODULE = $(shell go list -m)

.PHONY: generate build start migrate test lint
generate:
	go generate ./...

build: # build a server
	go build -a -o rena-server $(MODULE)/cmd/server

start: # start a server
	./rena-server server

factory: # factory
	./rena-server factory

migrate: # migrate database
	migrate -path ./migrations -database "postgres://postgres:password@nft-market.cqtyaecn0tih.us-east-2.rds.amazonaws.com:5432/postgres?sslmode=disable" up

test:
	go clean -testcache
	go test ./... -v

lint:
	gofmt -l .