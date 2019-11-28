
MAIN_PATH=$(shell pwd)
NAME=app
BIN_DIR=$(MAIN_PATH)/bin

all: build

build:	
	CGO_ENABLED=0 go build -v -o  $(BIN_DIR)/$(NAME) $(MAIN_PATH)/src

run:
	$(BIN_DIR)/$(NAME)

test:
	docker run -d --name testdb -e POSTGRES_USER=test_task_testdb -p "5432:5432" postgres:11
	@sleep 5
	docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://test_task_testdb:postgres@localhost:5432/test_task_testdb?sslmode=disable up 2
	- GO_ENABLED=0 go test -v $(MAIN_PATH)/src/db
	docker stop testdb && docker rm testdb

generate:
	swagger generate server --target $(MAIN_PATH)/src --exclude-main -A TestTask

up:
	docker-compose up

migrate-up:
	docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://test_task:postgres@localhost:5432/test_task?sslmode=disable up $(seq)

migrate-down:
	docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://test_task:postgres@localhost:5432/test_task?sslmode=disable down $(seq)
