# protège les appels aux commandes make xxx
.PHONY: kill build down

export DOCKER_PROJECT = ${PROJECT_NAME}

kill:
	docker compose kill

build:
	docker compose up -d --build

up:
	docker compose up -d

down:
	docker compose down