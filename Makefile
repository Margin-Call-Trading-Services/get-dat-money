detach ?= false

.PHONY: server kill test
server:
	./run.sh $(detach)

kill:
	docker compose down

test:
	go test ./... -cover
