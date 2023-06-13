.PHONY: help

help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

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
	goimports -d $(find . -type f -regexp '*.go' -not -path "./vendor/*")

.PHONY: vet
vet: ## Run vet.
	go vet -v ./...

.PHONY: staticcheck
staticcheck: ## Run staticcheck.
	staticcheck ./...

.PHONY: update-all-mod
update-all-mod: ## Run go get and mod.
	go get -u && go mod tidy && go get -u ./...

.PHONY: test-benchmark
test-benchmark: ## Run benchmark tests.
	go test -bench=. -benchmem

.PHONY: test-benchmark-static
test-benchmark-static: ## Run benchmark tests only static.
	go test -bench=^BenchmarkStatic.* -benchmem

.PHONY: test-benchmark-pathparam
test-benchmark-pathparam: ## Run benchmark tests only pathparam.
	go test -bench=^BenchmarkPathParam.* -benchmem

.PHONY: test-benchmark-by-regexp
test-benchmark-by-regexp: ## Run benchmark tests using regexp. ex. make test-benchmark-by-regexp EXP=Goblin, make test-benchmark-by-regexp EXP=StaticRoutes1
	go test -bench=.*${EXP}.* -benchmem

.PHONY: report-static-routes-root
report-static-routes-root: ## Run benchmark tests for reporing of StaticRoutesRoot.
	make test-benchmark-by-regexp EXP=StaticRoutesRoot | grep --line-buffered ^Bench > reports/tmp.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$2}' > reports/static-routes-root/time.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$3}' > reports/static-routes-root/nsop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$5}' > reports/static-routes-root/bop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$7}' > reports/static-routes-root/allocs.csv
	rm -rf reports/tmp.csv

.PHONY: report-static-routes-1
report-static-routes-1: ## Run benchmark tests for reporing of StaticRoutes1.
	make test-benchmark-by-regexp EXP=StaticRoutes1[^0] | grep --line-buffered ^Bench > reports/tmp.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$2}' > reports/static-routes-1/time.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$3}' > reports/static-routes-1/nsop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$5}' > reports/static-routes-1/bop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$7}' > reports/static-routes-1/allocs.csv
	rm -rf reports/tmp.csv

.PHONY: report-static-routes-5
report-static-routes-5: ## Run benchmark tests for reporing of StaticRoutes5.
	make test-benchmark-by-regexp EXP=StaticRoutes5 | grep --line-buffered ^Bench > reports/tmp.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$2}' > reports/static-routes-5/time.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$3}' > reports/static-routes-5/nsop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$5}' > reports/static-routes-5/bop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$7}' > reports/static-routes-5/allocs.csv
	rm -rf reports/tmp.csv

.PHONY: report-static-routes-10
report-static-routes-10: ## Run benchmark tests for reporing of StaticRoutes10.
	make test-benchmark-by-regexp EXP=StaticRoutes10 | grep --line-buffered ^Bench > reports/tmp.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$2}' > reports/static-routes-10/time.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$3}' > reports/static-routes-10/nsop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$5}' > reports/static-routes-10/bop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$7}' > reports/static-routes-10/allocs.csv
	rm -rf reports/tmp.csv

.PHONY: report-pathparam-routes-1
report-pathparam-routes-1: ## Run benchmark tests for reporing of PathParamRoutes1.
	make test-benchmark-by-regexp EXP=PathParamRoutes1[^0] | grep --line-buffered ^Bench > reports/tmp.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$2}' > reports/pathparam-routes-1/time.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$3}' > reports/pathparam-routes-1/nsop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$5}' > reports/pathparam-routes-1/bop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$7}' > reports/pathparam-routes-1/allocs.csv
	rm -rf reports/tmp.csv

.PHONY: report-pathparam-routes-5
report-pathparam-routes-5: ## Run benchmark tests for reporing of PathParamRoutes5.
	make test-benchmark-by-regexp EXP=PathParamRoutes5 | grep --line-buffered ^Bench > reports/tmp.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$2}' > reports/pathparam-routes-5/time.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$3}' > reports/pathparam-routes-5/nsop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$5}' > reports/pathparam-routes-5/bop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$7}' > reports/pathparam-routes-5/allocs.csv
	rm -rf reports/tmp.csv

.PHONY: report-pathparam-routes-10
report-pathparam-routes-10: ## Run benchmark tests for reporing of PathParamRoutes10.
	make test-benchmark-by-regexp EXP=PathParamRoutes10 | grep --line-buffered ^Bench > reports/tmp.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$2}' > reports/pathparam-routes-10/time.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$3}' > reports/pathparam-routes-10/nsop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$5}' > reports/pathparam-routes-10/bop.csv
	cat reports/tmp.csv | awk -v OFS=, '{print $$1,$$7}' > reports/pathparam-routes-10/allocs.csv
	rm -rf reports/tmp.csv

.PHONY: report
report: ## Run benchmark tests for reporing.
	make report-static-routes-root
	make report-static-routes-1
	make report-static-routes-5
	make report-static-routes-10
	make report-pathparam-routes-1
	make report-pathparam-routes-5
	make report-pathparam-routes-10
