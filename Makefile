.PHONY: run, start-front, build-front

run:
	docker-compose up

build-front:
	cd ui/registration-form/ && npm run build-dev

rebuild-backend:
	docker-compose build web
