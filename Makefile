docker_run:
	docker-compose up

updateFormUi:
	cd ui/registration-form/ && npm run build-dev

workOnFront:
	cd ui/registration-form/ && npm run start

run:
	go run cmd/main.go

.PHONY: docker_run form_ui run