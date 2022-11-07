build:
	docker build \
		-f dockerfile \
		-t api:1.0 \
		.

start:
	docker-compose up

stop:
	docker-compose down

tests:
	go test ./... -v -cover
