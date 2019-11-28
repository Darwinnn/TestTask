
MAIN_PATH=$(shell pwd)
NAME=app
BIN_DIR=$(MAIN_PATH)/bin

all: build

build:	
	CGO_ENABLED=0 go build -v -o  $(BIN_DIR)/$(NAME) $(MAIN_PATH)/src

run:
	$(BIN_DIR)/$(NAME)

test:
	GO_ENABLED=0 go test -v $(MAIN_PATH)/src

generate:
	swagger generate server --target $(MAIN_PATH)/src --exclude-main -A TestTask

up:
	docker-compose up

migrate-up:
	docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://test_task:postgres@localhost:5432/test_task?sslmode=disable up $(seq)

migrate-down:
	docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://test_task:postgres@localhost:5432/test_task?sslmode=disable down $(seq)
