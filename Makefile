go_entrypoint_path="./cmd/app/main.go"

db_url = "sqlite://$(DB_PATH)"
# title of migration
title = "migration"
version = 1

# --- #
# APP #
# --- #

dev:
	go run $(go_entrypoint_path)

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
migrate-force: check-db-env
	migrate -database $(db_url) -path migration -verbose force $(version)

migrate-version: check-db-env
	migrate -database $(db_url) -path migration -verbose version

migrate-up: check-db-env
	migrate -database $(db_url) -path migration -verbose up 1

migrate-down: check-db-env
	migrate -database $(db_url) -path migration -verbose down 1
