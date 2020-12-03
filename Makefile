SHELL=/bin/bash
# Go aliases
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_TOOL_COVER=$(GO_CMD) tool cover
GO_GET=$(GO_CMD) get
BINARY_NAME=godb
CMD_PATH=./cmd/main.go
BIN=bin
DIST=dist
DIST_MAC=$(DIST)/darwin
DIST_LINUX=$(DIST)/linux
DIST_WIN=$(DIST)/windows
DATABASE_URL='postgres://godb:godb@0.0.0.0:5432/godb?sslmode=disable'

build-linux:
	mkdir -p $(DIST_LINUX)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o ./$(DIST_LINUX)/$(BINARY_NAME) -v $(CMD_PATH)

build-mac:
	mkdir -p $(DIST_MAC)
	GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o ./$(DIST_MAC)/$(BINARY_NAME) -v $(CMD_PATH)

build-windows:
	mkdir -p $(DIST_WIN)
	GOOS=windows GOARCH=amd64 $(GO_BUILD) -o ./$(DIST_WIN)/$(BINARY_NAME).exe -v $(CMD_PATH)

build: build-linux build-mac build-windows

clean:
	rm -rf $(DIST)
	rm -rf $(BIN)

up:
	docker-compose up -d
	sleep 6
	go-migrate -database ${DATABASE_URL} -path db/migrations up
	DATABASE_URL=${DATABASE_URL} ./$(DIST_LINUX)/$(BINARY_NAME)

down:
	docker-compose down