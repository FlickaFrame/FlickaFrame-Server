GO ?= go
GOLANGCI_LINT_PACKAGE ?= github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1
GOPROXY ?= GOPROXY=https://goproxy.cn/,direct

# 伪目标，用于执行 go mod tidy
tidy:
	$(GOPROXY) go mod tidy

.PHONY: gen-web-api
gen-web-api: ## generate web api
	goctl api go -api ./spec/web/api/main.api -dir ./app/web/api --style go_zero
	make tidy

.PHONY: run-web
run-web: ## run web
	go run ./app/web/api/main.go -f ./app/web/api/etc/main.yaml

.PHONY: gen-user-rpc
gen-user-rpc: ## generate user rpc
	mkdir -p ./app/user/rpc
	cp ./spec/proto/user.proto ./app/user/rpc/
	cd ./app/user/rpc && goctl rpc protoc user.proto --go_out=pb --go-grpc_out=pb --zrpc_out=. --style go_zero
	rm -rf ./app/user/rpc/user.proto
	make tidy

.PHONY: gen-oss-rpc
gen-oss-rpc:
	mkdir -p ./app/oss/rpc
	cp ./spec/proto/oss.proto ./app/oss/rpc/
	cd ./app/oss/rpc && goctl rpc protoc oss.proto --go_out=pb --go-grpc_out=pb --zrpc_out=. --style go_zero
	rm -rf ./app/oss/rpc/oss.proto
	make tidy

.PHONY: run-oss-rpc
run-oss-rpc:
	go run app/oss/rpc/oss.go -f etc/oss_rpc.yml

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
