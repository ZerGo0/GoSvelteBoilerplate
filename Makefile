.PHONY: help

help:
	@echo 'Usage: make <target> [options]'
	@echo
	@echo 'Targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: start-server start-client

start-server: ## Usage: make start-server
	@go run .

start-client: ## Usage: make start-client
	@cd app && pnpm dev
