GO ?= go
GOLANGCI_LINT_PACKAGE ?= github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1

.PHONY: tidy
tidy: ## go mod tidy
	${GO} mod tidy

.PHONY: gen-api-go
gen-api-go: ## generate api go
	goctl api go --dir=./ --api ./desc/main.api  --style go_zero

.PHONY: gen-api-doc
gen-api-doc: ## generate api doc
	rm -rf docs/api/main.md
	goctl api doc --dir ./desc/  -o docs/api/


.PHONY: lint-go
lint-go:
	$(GO) run $(GOLANGCI_LINT_PACKAGE) run

.PHONY: lint-go-fix
lint-go-fix:
	$(GO) run $(GOLANGCI_LINT_PACKAGE) run --fix

.PHONY: build
build: ## build
	${GO} build -o tiktok .
