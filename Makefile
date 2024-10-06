GOCACHE=$(shell pwd)/.cache
export

ENV_FILE := $(strip $(wildcard ../.env.$(ENVIRONMENT)))
ifneq ($(ENV_FILE),)
	include $(ENV_FILE)
	export
endif

.PHONY: test

build:
	env
	go build -o rest .

generate:
	go run github.com/steebchen/prisma-client-go generate

migrate:
	@echo "Using environment file: $(ENV_FILE)"
	echo $(DB_URI)
	cd database && go run github.com/steebchen/prisma-client-go db push

server:
	go run .

server/debug:
	LOG_QUERIES=1 PHOTON_GO_LOG=info go run .
	go run .

seed/rbac:
	go run ./cmd/seed/rbac/main.go

test:
	cd locations && go test -v
	cd cameras && go test -v
	cd invitations && go test -v