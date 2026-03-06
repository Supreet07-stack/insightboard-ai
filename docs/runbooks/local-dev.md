# Local Development

## Start infrastructure

make up

## Stop infrastructure

make down

## Show running containers

make ps

## Show logs

make logs

## Infrastructure ports

- Postgres: localhost:5432
- Redis: localhost:6379
- RabbitMQ: localhost:5672
- RabbitMQ UI: http://localhost:15672
- OpenSearch: http://localhost:9200

## Credentials

### Postgres

- user: insight
- password: insight
- database: insightboard

### RabbitMQ

- user: insight
- password: insight

## Quick checks

### Postgres

```bash
docker exec -it insightboard-ai-postgres psql -U insight -d insightboard -c "SELECT now();"
```
