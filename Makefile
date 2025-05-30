# 颜色定义
COLOR_RESET = \033[0m
COLOR_BOLD_CYAN = \033[1;36m
COLOR_BOLD_GREEN = \033[1;32m
COLOR_YELLOW = \033[1;33m

# 定义项目相关变量
BINARY_NAME = test/main.wasm
BINARY_NAME_TINY = test/main_tiny.wasm
SRC_MAIN = main.go

.PHONY: serve build build-tiny clean deps test help

dev: serve

start:
	@$(MAKE) serve

# 调试页面
serve: build build-tiny
	@cd test && http-server -o index.html

# 构建go的wasm文件
build:
	@echo "$(COLOR_BOLD_CYAN)正在构建go的wasm文件...$(COLOR_RESET)"
	@GOOS=js GOARCH=wasm go build -o $(BINARY_NAME) $(SRC_MAIN)
	@echo "$(COLOR_BOLD_GREEN)构建完成!$(COLOR_RESET)"

# 构建tinygo的wasm文件
build-tiny:
	@echo "$(COLOR_BOLD_CYAN)正在构建tinygo的wasm文件...$(COLOR_RESET)"
	@tinygo build -o ./test/main_tiny.wasm -target=wasm main.go
	@echo "$(COLOR_BOLD_GREEN)构建完成!$(COLOR_RESET)"

# 清理构建文件
clean:
	@echo "$(COLOR_BOLD_CYAN)正在清理构建文件...$(COLOR_RESET)"
	@rm -f ./test/*.wasm
	@echo "$(COLOR_BOLD_GREEN)清理完成!$(COLOR_RESET)"

# 安装依赖
deps:
	@echo "$(COLOR_BOLD_CYAN)正在安装依赖...$(COLOR_RESET)"
	@go mod tidy
	@echo "$(COLOR_BOLD_GREEN)依赖安装完成!$(COLOR_RESET)"

# 运行测试
test:
	@echo "$(COLOR_BOLD_CYAN)正在运行测试...$(COLOR_RESET)"
	@go test -v ./...
	@echo "$(COLOR_BOLD_GREEN)测试完成!$(COLOR_RESET)"

# 显示帮助信息
help:
	@echo "$(COLOR_BOLD_CYAN)可用命令:$(COLOR_RESET)"
	@echo "  $(COLOR_BOLD_GREEN)make build$(COLOR_RESET)          - 构建wasm文件"
	@echo "  $(COLOR_BOLD_GREEN)make make build-tiny$(COLOR_RESET) - 构建tinygo的wasm文件"
	@echo "  $(COLOR_BOLD_GREEN)make clean$(COLOR_RESET)  - 清理构建文件"
	@echo "  $(COLOR_BOLD_GREEN)make deps$(COLOR_RESET)   - 安装依赖"
	@echo "  $(COLOR_BOLD_GREEN)make test$(COLOR_RESET)   - 运行测试"
	@echo "  $(COLOR_YELLOW)make help$(COLOR_RESET)   - 显示此帮助信息"
