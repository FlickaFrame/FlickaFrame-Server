GO ?= go
GOLANGCI_LINT_PACKAGE ?= github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1

.PHONY: install
install: ## install
	GOPROXY=https://goproxy.cn/,direct ${GO} install github.com/zeromicro/go-zero/tools/goctl@latest

.PHONY: tidy
tidy: ## go mod tidy
	GOPROXY=https://goproxy.cn/,direct ${GO} mod tidy

.PHONY: gen-api-go
gen-api-go: ## generate api go
	goctl api go --dir=./ --api ./desc/main.api  --style go_zero

.PHONY: gen-api-doc
gen-api-doc: ## generate api doc
	rm -rf docs/api/*
	mkdir -p docs/api/
	goctl api doc --dir ./desc/  -o docs/api/

.PHONY: api-format
api-format: ## api format
	goctl api format --dir ./desc/

.PHONY: gen-api
gen-api: gen-api-go gen-api-doc api-format ## generate api

.PHONY: build
build: tidy
	${GO} build -o main .

.PHONY: run
run: gen-api tidy
	${GO} run main.go
