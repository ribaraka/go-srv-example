.PHONY: run, build-front, rebuild-backend

run:
	docker-compose up

build-front:
	cd ui/registration-form/ && npm run build-dev

rebuild-backend:
	docker-compose build web
