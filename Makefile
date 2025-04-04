up-build:
	docker compose up --build -d
app-rebuild:
	docker compose build auth
	docker compose up --no-deps auth
down:
	docker compose down
local-run:
	CONFIG_PATH='./configs/dev.yml' go run cmd/main.go