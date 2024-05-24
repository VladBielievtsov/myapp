run: build
	@./bin/server

build:
	@go build -o bin/server cmd/api/main.go

seed:
	@go run cmd/seed/main.go

drop:
	@go run cmd/drop/main.go