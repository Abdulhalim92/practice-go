rabbitmq:
	docker compose -f rabbitmq/rabbit-compose.yaml up -d

stop:
	docker compose -f rabbitmq/rabbit-compose.yaml down

.PHONY: rabbitmq stop