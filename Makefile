run:
	@go run cmd/main.go

generate:
	@buf generate
lint:
	@golangci-lint run