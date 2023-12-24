GO ?= go
GOLANGCI_LINT_PACKAGE ?= github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1
GOPROXY ?= GOPROXY=https://goproxy.cn/,direct

# 伪目标，用于执行 go mod tidy
tidy:
	$(GOPROXY) go mod tidy

.PHONY: install
install:
	GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/zeromicro/go-zero/tools/goctl
	GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-go-compact@latest

.PHONY: gen
	goctl api go -api

 	## generate Web-API
	goctl api plugin -p goctl-go-compact -api ./spec/web/api/main.api -dir ./app/web/api --style go_zero
	goctl api plugin -p goctl-go-compact -api ./spec/web/api/oss_server.api -dir ./app/oss/api --style go_zero

	## generate User-RPC
	mkdir -p ./app/user/rpc
	cp ./spec/proto/user.proto ./app/user/rpc/
	cd ./app/user/rpc && goctl rpc protoc user.proto --go_out=pb --go-grpc_out=pb --zrpc_out=. --style go_zero
	rm -rf ./app/user/rpc/user.proto

	## generate Oss-RPC
	mkdir -p ./app/oss/rpc
	cp ./spec/proto/oss.proto ./app/oss/rpc/
	cd ./app/oss/rpc && goctl rpc protoc oss.proto --go_out=pb --go-grpc_out=pb --zrpc_out=. --style go_zero
	rm -rf ./app/oss/rpc/oss.proto

	## generate Follow-RPC
	mkdir -p ./app/follow/rpc
	cp ./spec/proto/follow.proto ./app/follow/rpc/
	cd ./app/follow/rpc && goctl rpc protoc follow.proto --go_out=pb --go-grpc_out=pb --zrpc_out=. --style go_zero
	rm -rf ./app/follow/rpc/follow.proto

	## generate Favorite-RPC
	mkdir -p ./app/favorite/rpc
	cp ./spec/proto/favorite.proto ./app/favorite/rpc/
	cd ./app/favorite/rpc && goctl rpc protoc favorite.proto --go_out=pb --go-grpc_out=pb --zrpc_out=. --style go_zero
	rm -rf ./app/favorite/rpc/favorite.proto
	make tidy

.PHONY: run-oss-rpc
run-oss-rpc:
	go run app/oss/rpc/oss.go -f etc/oss_rpc.yml


.PHONY: run-follow-rpc
run-follow-rpc:
	go run app/follow/rpc/follow.go -f etc/follow_rpc.yml



.PHONY: build
build: tidy
	${GO} build -o main .

.PHONY: run
run: gen-api tidy
	${GO} run main.go
