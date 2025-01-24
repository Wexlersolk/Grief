include .envrc
MIGRATIONS_PATH = ./cmd/migrate/migrations

# Test
.PHONY: test
test:
	@go test -v ./...

# Application Database Container
create_db_container:
	@echo "Creating Application PostgreSQL container..."
	docker run --name ${DB_DOCKER_CONTAINER} -p ${DB_PORT}:5432 -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -e POSTGRES_DB=${DB_NAME} -d postgres:12-alpine

create_app_db:
	@echo "Creating application database ${DB_NAME}..."
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

# Keycloak Database Container
create_userdb_container:
	@echo "Creating Keycloak PostgreSQL container..."
	docker run --name ${USERDB_DOCKER_CONTAINER} -p ${USERDB_PORT}:5432 -e POSTGRES_USER=${USERDB_USER} -e POSTGRES_PASSWORD=${USERDB_PASSWORD} -e POSTGRES_DB=${USERDB_NAME} -d postgres:12-alpine

create_userdb:
	@echo "Creating Keycloak database ${USERDB_NAME}..."
	docker exec -it ${USERDB_DOCKER_CONTAINER} createdb --username=${USERDB_USER} --owner=${USERDB_USER} ${USERDB_NAME}

# Keycloak Container
create_keycloak_container:
	@echo "Creating Keycloak container..."
	docker run -d --name ${KEYCLOAK_CONTAINER} -p ${KEYCLOAK_PORT}:8080 -e KEYCLOAK_ADMIN=${KEYCLOAK_ADMIN_USER} -e KEYCLOAK_ADMIN_PASSWORD=${KEYCLOAK_ADMIN_PASSWORD} -e KC_DB=postgres -e KC_DB_URL=jdbc:postgresql://${USERDB_DOCKER_CONTAINER}:5432/${USERDB_NAME} -e KC_DB_USERNAME=${USERDB_USER} -e KC_DB_PASSWORD=${USERDB_PASSWORD} --link ${USERDB_DOCKER_CONTAINER}:postgres quay.io/keycloak/keycloak:latest start-dev

# Start Containers
start_db_container:
	@echo "Starting Application PostgreSQL container..."
	docker start ${DB_DOCKER_CONTAINER}

start_userdb_container:
	@echo "Starting Keycloak PostgreSQL container..."
	docker start ${USERDB_DOCKER_CONTAINER}

start_keycloak_container:
	@echo "Starting Keycloak container..."
	docker start ${KEYCLOAK_CONTAINER}

start_containers: start_db_container start_userdb_container start_keycloak_container
	@echo "Started all containers."

# Stop Containers
stop_db_container:
	@echo "Stopping Application PostgreSQL container..."
	docker stop ${DB_DOCKER_CONTAINER}

stop_userdb_container:
	@echo "Stopping Keycloak PostgreSQL container..."
	docker stop ${USERDB_DOCKER_CONTAINER}

stop_keycloak_container:
	@echo "Stopping Keycloak container..."
	docker stop ${KEYCLOAK_CONTAINER}

stop_containers: stop_db_container stop_userdb_container stop_keycloak_container
	@echo "Stopped all containers."

# Remove Containers
remove_db_container:
	@echo "Removing Application PostgreSQL container..."
	docker rm -f ${DB_DOCKER_CONTAINER}

remove_userdb_container:
	@echo "Removing Keycloak PostgreSQL container..."
	docker rm -f ${USERDB_DOCKER_CONTAINER}

remove_keycloak_container:
	@echo "Removing Keycloak container..."
	docker rm -f ${KEYCLOAK_CONTAINER}

remove_containers: remove_db_container remove_userdb_container remove_keycloak_container
	@echo "Removed all containers."

# Migrations
.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

# Seed
.PHONY: seed
seed: 
	@go run cmd/migrate/seed/main.go

# Generate Docs
.PHONY: gen-docs
gen-docs:
	@swag init -g ./api/main.go -d cmd,internal && swag fmt
