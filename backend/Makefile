.PHONY: build
build:
	docker compose build

.PHONY: up
up:
	docker compose up --build

.PHONY: down
down:
	docker compose down -v

.PHONY: lint
lint:
	docker compose --profile extra run buf buf lint

.PHONY: gen
gen:
	docker compose --profile extra run buf buf generate