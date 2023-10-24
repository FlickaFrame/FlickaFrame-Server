GO ?= go

.PHONY: tidy
tidy: ## go mod tidy
	${GO} mod tidy

.PHONY: gen-api
gen-api: ## generate api
	goctl api go --dir=./ --api ./desc/main.api  --style go_zero
	${GO} mod tidy