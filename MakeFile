export
TESTABLE=$$(go list ./...)

run:
	@go run ./src/main/main.go $(day)
.PHONY: run

test:
	@go test ./internal/challenges/day-$(day)
.PHONY: test

test-all:
	@go test $(TESTABLE)
.PHONY: test-all