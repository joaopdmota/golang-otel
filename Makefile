.PHONY: up
up:
	@echo "Initializing..."
	docker compose up -d
	@echo "Initialization finished."

.PHONY: test
test:
	@echo "Running tests..."
	go test $(shell go list ./... | grep -v '/mocks') -coverprofile=coverage.out
	go tool cover -func=coverage.out
	@echo "Tests completed."


.PHONY: test-cover
test-cover:
	@echo "Running tests..."
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Tests completed."