build:
	go build

run:
	go run main.go

migration:
	docker run -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=go_project -d -p "33333":"5432" --name database postgres:latest

createDB:
	PGPASSWORD=password psql -h localhost -p 33333 -U postgres -d go_project -a -f ./migration/users_table.sql
	PGPASSWORD=password psql -h localhost -p 33333 -U postgres -d go_project -a -f ./migration/credentials.sql
	PGPASSWORD=password psql -h localhost -p 33333 -U postgres -d go_project -a -f ./migration/email_verification_tokens.sql

install:
	npm install

dockbuild:
 dockrun:
	docker run -p 8081:8081 --rm -v /home/ribaraka/projects/src/github.com/ribaraka/go-srv-example/cmd/:/config server:latest


.PHONY: docker install createDB migration build run dockrun