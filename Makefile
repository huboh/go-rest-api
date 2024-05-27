APP_PATH=./
APP_BIN_PATH=./bin/main

all: build run

run:
	go run ${APP_PATH}

build:
	GOOS=linux go build -o ${APP_BIN_PATH} ${APP_PATH}