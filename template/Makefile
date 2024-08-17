$(shell cp -n .env.example .env)
include .env
export

all: run
.PHONY: all

run: dbup
	@air
.PHONY: run

test: 
	@go test -v ./...
.PHONY: test

dbup:
	@docker compose -f docker/postgres/docker-compose.yml up -d
.PHONY: dbup

dbdown:
	@docker compose -f docker/postgres/docker-compose.yml down
.PHONY: dbdown

gendata: dbdown dbup
	@go test -v -tags=gendata -count=1 ./...
.PHONY: gendata

air: 
	@echo "Installing air"
	@go install github.com/air-verse/air@latest
.PHONY: air

templ:
	@echo "Installing templ"
	@rm -f $$(which templ) && go install github.com/a-h/templ/cmd/templ@$$(go list -m github.com/a-h/templ | cut -d' ' -f2)
.PHONY: templ

alpine:
	@echo "Installing alpine.js"
	@cd internal/server/assets/js && { curl -O https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js ; mv cdn.min.js alpine.js ; cd - ; } 
.PHONY: alpine

htmx:
	@echo "Installing htmx"
	@cd internal/server/assets/js && { curl -O https://unpkg.com/htmx.org@2.0.1/dist/htmx.min.js ; curl -O https://unpkg.com/htmx-ext-response-targets@2.0.0/response-targets.js ; cd - ; }
.PHONY: htmx

deps: air templ htmx alpine
	@go mod tidy
.PHONY: deps
