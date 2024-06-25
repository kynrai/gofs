$(shell cp -n .env.example .env)
include .env
export


all: air
.PHONY: all

run:
	@go run cmd/server/main.go
.PHONY: run

test: 
	@go test -v ./...
.PHONY: test

air:
	@air
.PHONY: air

templ:
	@rm -f $$(which templ) && go install github.com/a-h/templ/cmd/templ@$$(go list -m github.com/a-h/templ | cut -d' ' -f2)
.PHONY: templ
