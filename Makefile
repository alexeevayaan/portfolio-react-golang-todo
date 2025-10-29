run:
	go run ./api/cmd/app

up:
	docker compose up --build --force-recreate

down:
	docker compose down

