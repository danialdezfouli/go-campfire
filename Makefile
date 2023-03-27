.PHONY: all clean

APPNAME := myapp
GOFILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

all: build run

build:
	go build -o $(APPNAME) cmd/main.go

run: 
	go run cmd/web/main.go

up: 
	docker compose --env-file .env --file ./infrastructure/docker-compose.yml up

migrate: 
	go run cmd/migrate/main.go

clean:
	rm -f $(APPNAME)