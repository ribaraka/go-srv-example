.PHONY: build
build:
	go build
# create docker file
.PHONY: run
	go run main.go
.DEFAULT_GOAL := run

.PHONY: migration
migration:
	docker run -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=go_project -d -p "33333":"5432" postgres

.PHONY: createDB
createDB:
	PGPASSWORD=password psql -h localhost -p 33333 -U postgres -d go_project -a -f ./migration/credentials.sql
	PGPASSWORD=password psql -h localhost -p 33333 -U postgres -d go_project -a -f ./migration/email_verification_tokens.sql
	PGPASSWORD=password psql -h localhost -p 33333 -U postgres -d go_project -a -f ./migration/users_table.sql

.PHONY: install
install:
	npm install