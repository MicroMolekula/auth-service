up-build:
	docker compose up --build -d
app-rebuild:
	docker compose build app
	docker compose up --no-deps app
local-run:
	CONFIG_PATH='./configs/dev.yml' go run cmd/main.go