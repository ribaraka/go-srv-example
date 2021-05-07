database:
	docker run -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=go_project -d -p "33333":"5432" --name database postgres:latest

create-table:
	PGPASSWORD=password psql -h localhost -p 33333 -U postgres -d go_project -a -f ./migration/users_table.sql
	PGPASSWORD=password psql -h localhost -p 33333 -U postgres -d go_project -a -f ./migration/credentials.sql
	PGPASSWORD=password psql -h localhost -p 33333 -U postgres -d go_project -a -f ./migration/email_verification_tokens.sql

dev-build:
	docker build -t server .

dev-run:
	docker run -p 8081:8081 --rm --net=host -v /home/ribaraka/projects/src/github.com/ribaraka/go-srv-example/cmd/:/config server:latest -confile ./config/config.yaml

form_ui:
	npm install

.PHONY: database create-table dev-build dev-run