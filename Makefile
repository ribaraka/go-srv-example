docker_run:
	docker-compose up

buildFront:
	cd ui/registration-form/ && npm run build-dev

launchFront:
	cd ui/registration-form/ && npm run start

run:
	go run cmd/main.go

.PHONY: docker_run form_ui run