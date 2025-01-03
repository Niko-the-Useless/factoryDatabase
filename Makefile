BINARY_NAME=factoryDbAPI
GOFILES= server.go

all: run

build:
	go build -o $(BINARY_NAME) $(GOFILES)

run:
	go run $(GOFILES)

test:
	go test

clean:
	rm -f $(BINARY_NAME)
