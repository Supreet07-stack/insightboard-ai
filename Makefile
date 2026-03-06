.PHONY: up down logs ps fmt

up:
	docker compose up -d

down:
	docker compose down -v

logs:
	docker compose logs -f --tail=200

ps:
	docker compose ps

fmt:
	find services libs -name "*.go" -print0 | xargs -0 gofmt -w