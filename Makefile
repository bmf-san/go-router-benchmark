.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

.PHONY: install-tools
install-tools: ## Install staticcheck.
ifeq ($(shell command -v staticcheck 2> /dev/null),)
	go install honnef.co/go/tools/cmd/staticcheck@latest
endif
ifeq ($(shell command -v goimports 2> /dev/null),)
	go install golang.org/x/tools/cmd/goimports@latest
endif

.PHONY: gofmt
gofmt: ## Run gofmt.
	test -z "$(gofmt -s -l . | tee /dev/stderr)"

.PHONY: goimports
goimports: ## Run goimports.
	goimports -d $(find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: vet
vet: ## Run vet.
	go vet -v ./...

.PHONY: staticcheck
staticcheck: ## Run staticcheck.
	staticcheck ./...

.PHONY: test-benchmark
test-benchmark: ## Run benchmark tests.
	go test -bench=. -benchmem