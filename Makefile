detach ?= false

server:
	./scripts/server.sh $(detach)

kill:
	docker compose down