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
	rm -rf docs/api/*
	goctl api doc --dir ./desc/  -o docs/api/

.PHONY: gen-api-swagger
gen-api-swagger: ## generate api swagger
	GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest
	rm -rf docs/swagger/*
	goctl api plugin -plugin goctl-swagger="swagger -filename main.json" -api desc/main.api -dir docs/swagger

.PHONY: api-format
api-format: ## api format
	goctl api format --dir ./desc/

.PHONY: lint-go
lint-go:
	$(GO) run $(GOLANGCI_LINT_PACKAGE) run

.PHONY: lint-go-fix
lint-go-fix:
	$(GO) run $(GOLANGCI_LINT_PACKAGE) run --fix

.PHONY: build
build: ## build
	GOPROXY=https://goproxy.cn/,direct go mod tidy
	${GO} build -o main .
