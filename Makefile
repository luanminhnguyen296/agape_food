.PHONY: gen run build clean help

.DEFAULT_GOAL := help

help: ## Show available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

gen: ## Generate templ templates
	templ generate

run: gen ## Run the server locally with live reload/generation
	go run main.go

build: gen ## Build the production binary
	GOOS=darwin GOARCH=arm64 go build -o landingpage main.go

static: gen ## Build static HTML and assets into dist/ for static hosting
	go run cmd/build_static/main.go

clean: ## Clean generated templ files, binary, and static dist
	rm -f landingpage
	rm -rf dist
	find . -name "*_templ.go" -type f -delete