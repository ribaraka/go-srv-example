docker_run:
	docker-compose up

form_ui:
	cd ui/registrationForm/ && npm run build

run:
	go run cmd/main.go

.PHONY: docker_run form_ui run