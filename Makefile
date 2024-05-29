BINARY_NAME = clocky

CMD_DIR = ./cmd/clocky

INSTALL_PATH = $(GOPATH)/bin/$(BINARY_NAME)

.PHONY: all
all: build

.PHONY: build
build:
	# env GOOS=windows GOARCH=amd64 go build $(CMD_DIR)
	env GOOS=darwin GOARCH=arm64 go build $(CMD_DIR)
	mv clocky clockym
	go build $(CMD_DIR)

.PHONY: install
install:
	go install $(CMD_DIR)

.PHONY: clean
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)m
	rm -f $(BINARY_NAME).exe

.PHONY: run
run:
	go run $(CMD_DIR)

.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  all      - Default target, builds the binary"
	@echo "  build    - Build the binary"
	@echo "  install  - Install the binary to the GOPATH/bin directory"
	@echo "  clean    - Clean up binary and other generated files"
	@echo "  run      - Run the program"
	@echo "  help     - Show this help message"
