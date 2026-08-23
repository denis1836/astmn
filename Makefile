NAME=astmn
MAIN_PATH=./cmd/astmn

GO=go
LINTER=golangci-lint

.PHONY: all build run test lint fmt vet tidy clean help

all: build

build:
	@echo "> Compiling $(NAME)..."
	$(GO) build -o $(NAME) $(MAIN_PATH)

run: build
	./$(NAME) --help

test:
	@echo "> Running tests..."
	$(GO) test -v ./...

fmt:
	@echo "> Formatting code..."
	$(GO) fmt ./...

lint: fmt
	@echo "> Linting code..."
	$(LINTER) run ./...

vet:
	@echo "> Verifying code..."
	$(GO) vet ./...

tidy:
	@echo "> Cleaning dependencies..."
	$(GO) mod tidy

clean:
	@echo "> Cleaning artifacts..."
	$(GO) clean
	rm -f $(NAME)
	rm -rf .astmn/logs .astmn/project.db*
	rm -rf Mod/ Content/ /tmp/astmn

help:
	@echo "Available Makefile commands:"
	@grep -E '^[a-zA-Z_-]+:' $(MAKEFILE_LIST) | sed 's/://' | grep -v 'PHONY'