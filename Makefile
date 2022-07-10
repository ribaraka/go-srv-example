.PHONY: run
run:
	docker-compose up

.PHONY: stop
stop:
	docker-compose down

.PHONY: build-front
build-front:
	cd ui/registration-form/ && npm run build-dev

.PHONY: rebuild-backend
rebuild-backend:
	docker-compose build web
