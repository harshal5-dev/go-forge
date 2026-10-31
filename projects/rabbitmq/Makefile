.PHONY: rabbit-up rabbit-down

rabbit-up:
	docker run -d --name rabbitmq-local -p 5672:5672 -p 15672:15672 -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=12345 rabbitmq:4.4.0-management-alpine

rabbit-down:
	docker rm -f rabbitmq-local
