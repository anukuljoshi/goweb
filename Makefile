include .env

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo "Usage:"
	@sed -n "s/^##//p" ${MAKEFILE_LIST} | column -t -s ":" | sed -e "s/^/ /"

.PHONY: confirm
confirm:
	@echo -n "Are you sure? [y/N] " && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

.PHONY: run
## run: run the cmd/api application
run:
	@air

.PHONY: db/sql
## db/psql: connect to the database using psql
db/psql:
	psql ${DSN}

.PHONY: db/migrations/new
## db/migrations/new name=$1: create a new database migration
db/migrations/new:
	@echo "Creating migartion files for ${name}"
	migrate create -seq -ext=.sql -dir=./db/migrations ${name}

.PHONY: db/migrations/up
## db/migrations/up: apply all up database migrations
db/migrations/up: confirm
	@echo "Running up migrations"
	migrate -path ./db/migrations -database ${DSN} up

.PHONY: db/migrations/down
## db/migrations/down: apply all down database migrations
db/migrations/down: confirm
	@echo "Running down migrations"
	migrate -path ./db/migrations -database ${DSN} down

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

## fmt: format and tidy
.PHONY: fmt
fmt:
	@echo 'Formatting...'
	go fmt ./...
	@echo 'Tidying...'
	go mod tidy
	go mod verify

# ==================================================================================== #
# BUILD
# ==================================================================================== #

current_time = $(shell date --iso-8601=seconds)
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build: build the cmd/api application
.PHONY: build
build:
	@echo 'Building cmd/api...'
	go build -ldflags=${linker_flags} -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/api ./cmd/api
