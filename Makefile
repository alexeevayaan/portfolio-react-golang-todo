DB_MIGRATE_URL = postgres://login:pass@localhost:5432/postgres?sslmode=disable
MIGRATE_PATH = ./api/migration/postgres


run:
	go run ./api/cmd/app/main.go
# 	go run ./api/cmd/app

up:
	docker compose up --build --force-recreate

down:
	docker compose down

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

migrate-create:
	@read -p "Name:" name; \
	migrate create -ext sql -dir "$(MIGRATE_PATH)" $$name

migrate-up:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" up

migrate-down:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" down -all

migrate-drop:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" drop

migrate-version:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" version

integration-test:
	cd api && go test -count=1 -v -tags=integration ./test/integration