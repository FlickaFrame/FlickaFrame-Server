GO ?= go
GOLANGCI_LINT_PACKAGE ?= github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1
GOPROXY ?= GOPROXY=https://goproxy.cn/,direct

# 伪目标，用于执行 go mod tidy
tidy:
	$(GOPROXY) go mod tidy

.PHONY: install
install:
	GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero/tools/goctl

.PHONY: gen-web-api
gen-web-api: ## generate web api
	goctl api go -api ./spec/web/api/main.api -dir ./app/web/api --style go_zero
	make tidy

.PHONY: run-web-api
run-web-api: ## run web
	go run ./app/web/api/main.go -f ./etc/web_api.yml

.PHONY: gen-user-rpc
gen-user-rpc: ## generate user rpc
	mkdir -p ./app/user/rpc
	cp ./spec/proto/user.proto ./app/user/rpc/
	cd ./app/user/rpc && goctl rpc protoc user.proto --go_out=pb --go-grpc_out=pb --zrpc_out=. --style go_zero
	rm -rf ./app/user/rpc/user.proto
	make tidy

.PHONY: run-user-rpc
run-user-rpc:
	go run app/user/rpc/user.go -f etc/user_rpc.yml

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

.PHONY: ossRpc-dockerfile
ossRpc-dockerfile:
	cp ./etc/oss_rpc.yml ./app/oss/rpc/etc/oss.yaml
	rm -rf ./app/oss/rpc/Dockerfile
	cd ./app/oss/rpc && goctl docker -go oss.go  --port 8001

.PHONY: userRpc-dockerfile
userRpc-dockerfile:
	cp ./etc/user_rpc.yml ./app/user/rpc/etc/user.yaml
	rm -rf ./app/user/rpc/Dockerfile
	cd ./app/user/rpc && goctl docker -go user.go  --port 8002

.PHONY: followRpc-dockerfile
followRpc-dockerfile:
	cp ./etc/follow_rpc.yml ./app/follow/rpc/etc/follow.yaml
	rm -rf ./app/follow/rpc/Dockerfile
	cd ./app/follow/rpc && goctl docker -go follow.go  --port 8003

.PHONY: favoriteRpc-dockerfile
favoriteRpc-dockerfile:
	cp ./etc/favorite_rpc.yml ./app/favorite/rpc/etc/favorite.yaml
	rm -rf ./app/favorite/rpc/Dockerfile
	cd ./app/favorite/rpc && goctl docker -go favorite.go  --port 8004

.PHONY: webApi-dockerfile
webApi-dockerfile:
	cp ./etc/web_api.yml ./app/web/api/etc/main.yaml
	rm -rf ./app/web/api/Dockerfile
	cd ./app/web/api && goctl docker -go main.go  --port 8080

.PHONY: gen-dockerfile
gen-dockerfile: ossRpc-dockerfile userRpc-dockerfile followRpc-dockerfile favoriteRpc-dockerfile webApi-dockerfile

.PHONY: gen-follow-rpc
gen-follow-rpc:
	mkdir -p ./app/follow/rpc
	cp ./spec/proto/follow.proto ./app/follow/rpc/
	cd ./app/follow/rpc && goctl rpc protoc follow.proto --go_out=pb --go-grpc_out=pb --zrpc_out=. --style go_zero
	rm -rf ./app/follow/rpc/follow.proto
	make tidy

.PHONY: run-follow-rpc
run-follow-rpc:
	go run app/follow/rpc/follow.go -f etc/follow_rpc.yml

.PHONY: gen-favorite-rpc
gen-favorite-rpc:
	mkdir -p ./app/favorite/rpc
	cp ./spec/proto/favorite.proto ./app/favorite/rpc/
	cd ./app/favorite/rpc && goctl rpc protoc favorite.proto --go_out=pb --go-grpc_out=pb --zrpc_out=. --style go_zero
	rm -rf ./app/favorite/rpc/favorite.proto
	make tidy

.PHONY: run-favorite-rpc
run-favorite-rpc:
	go run app/favorite/rpc/favorite.go -f etc/favorite_rpc.yml

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
