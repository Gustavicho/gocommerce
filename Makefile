build:
		@go build -o bin/gocommerce cmd/main.go

test:
		@go test -v ./...

run: build
		@./bin/gocommerce