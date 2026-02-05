# Variables
APP_NAME=parking-api
COMPOSE=docker-compose

.PHONY: dev build up down restart logs refresh seed help

help:
	@echo "Usage: make [command]"
	@echo ""
	@echo "Commands:"
	@echo "  dev      : Rebuild, start containers, and refresh database"
	@echo "  up       : Start containers in background"
	@echo "  down     : Stop and remove containers"
	@echo "  refresh  : Run database refresh inside docker"
	@echo "  seed     : Run database seeder inside docker"
	@echo "  logs     : Follow application logs"

dev: build
	@clear
	@printf "\033[1;36m>>> INFRASTRUCTURE STATUS\033[0m\n"
	@printf "Grafana  : http://localhost:3000\n"
	@printf "Loki     : http://loki:3100\n"
	@printf "Metrics  : http://localhost:8080/metrics\n"
	@printf -- "----------------------------------------------------------\n"
	@printf "Waiting for container to be ready...\n"
	@sleep 5 # Kasih nafas 5 detik buat container startup
	@make refresh
	@printf -- "----------------------------------------------------------\n"
	@printf "\033[1;32m>>> SYSTEM IS LIVE\033[0m\n"
	@make logs

build:
	@printf "\033[1;36m>>> BUILDING DOCKER IMAGES\033[0m\n"
	$(COMPOSE) up -d --build

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

refresh:
	@printf "\033[1;36m>>> DATABASE REFRESH\033[0m\n"
	@docker exec -it $(APP_NAME) ./refresh-db

seed:
	@printf "\033[1;36m>>> SEEDING DATABASE\033[0m\n"
	docker exec -it -w /app $(APP_NAME) go run scripts/seed.go

logs:
	docker logs -f $(APP_NAME)