.PHONY: serve build build-tiny clean deps test help

dev: serve

start:
	$(MAKE) serve

serve: build build-tiny
	cd test && http-server -o index.html

build:
	GOOS=js GOARCH=wasm go build -o ./test/main.wasm main.go

build-tiny:
	tinygo build -o ./test/main_tiny.wasm -target=wasm main.go

clean:
	rm -f ./test/*.wasm

deps:
	go mod tidy

test:
	go test -v ./...

help:
	@echo "可用命令:"
	@echo "  make build  - 构建wasm文件"
	@echo "  make build-tiny  - 构建tinygo的wasm文件"
	@echo "  make clean  - 清理构建文件"
	@echo "  make deps   - 安装依赖"
	@echo "  make test   - 运行测试"
