go_app_path="./cmd/app/main.go"
go_migrator_path="./cmd/migrator/main.go"

db_url = "sqlite://$(DB_PATH)"
# title of migration
title = "migration"
version = 1

# --- #
# APP #
# --- #

dev:
	go run $(go_app_path)

lint:
	golangci-lint run -c ./.golangci.yml ./...

clean-lint-cache:
	golangci-lint cache clean

# ---------- #
# MIGRATIONS #
# ---------- #

check-db-env:
	@if [ -z "$(DB_PATH)" ]; then \
		echo "DB_PATH env is not presented"; \
		exit 2; \
	fi

# use "title" var for name the migration
migrations:
	migrate create -digits 2 -dir migration -ext sql -seq "$(title)"

# use "version" var for specify the version of the migration to force
migrate-force:
	@go run $(go_migrator_path) force $(version)

migrate-status:
	@go run $(go_migrator_path) status

migrate-up:
	@go run $(go_migrator_path) up -n 1

migrate-down:
	@go run $(go_migrator_path) down -n 1
