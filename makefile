test:
	@go test . -cover

coverage:
	@go test . -coverprofile=addr.coverage
	@go tool cover -html=addr.coverage
	@rm addr.coverage
