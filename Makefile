.PHONY: run

-include: .env

run:
	@kill -9 $$(lsof -ti :8080) 2>/dev/null || true
	@go run main.go