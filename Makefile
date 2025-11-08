run:
	go run main.go

tidy:
	go mod tidy

build:
	go build -o bin/server main.go

test:
	go test ./...

clean:
	rm -rf bin

migrate:
	go run migrate/migrate.go
