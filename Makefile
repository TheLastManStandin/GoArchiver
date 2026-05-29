GO := go
BINARY := moonArchiver
MAIN_PACKAGE := ./src

.PHONY: all test build clean

all: test build

test:
	$(GO) test ./...

build:
	$(GO) build -o $(BINARY) $(MAIN_PACKAGE)

clean:
	rm -f $(BINARY)
