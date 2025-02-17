build:
	go build -o ./bin/fadden ./main.go

run: build
	./bin/fadden

test:
	go test -v ./...
