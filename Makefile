docker_run:
	docker-compose up

buildFront:
	cd ui/registration-form/ && npm run build-dev

startFront:
	cd ui/registration-form/ && npm run start

goRun:
	go run cmd/main.go

.PHONY: docker_run form_ui run