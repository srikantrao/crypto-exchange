build:
	go build -o bin/exchange

run: build
	./bin/exchange

test:
	gotest -v ./...