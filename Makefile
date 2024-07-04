APP_NAME = easy
DOCKER_COMPOSE_FILE = docker-compose.yml

.PHONY: build
build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

.PHONY: up
up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

.PHONY: down
down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

.PHONY: logs
logs:
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

.PHONY: deploy
deploy: build up
