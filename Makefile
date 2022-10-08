ifndef $(GOROOT)
    GOROOT=$(shell go env GOROOT)
    export GOROOT
endif

include .env
export

DIR=$(notdir $(shell pwd))
export DIR

init:
	@echo -e "\n\t🔑\n"
	@go run $(GOROOT)/src/crypto/tls/generate_cert.go --host $(API_DOMAIN)

gen:
	@echo -e "\n\t🧠\n"
	@go run github.com/99designs/gqlgen generate
	sqlc generate

sql:
	@echo -e "\n\t🧠\n"
	@sqlc generate

dev:
	@echo -e "\n\t💣\n"
	docker-compose -p ${PROJECT} up --build --force-recreate --remove-orphans

.PHONY: init gen dev sql

