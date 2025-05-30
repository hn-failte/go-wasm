# go-wasm

Go WebAssembly 示例项目，演示如何使用 Go 和 TinyGo 构建 WebAssembly 模块并在浏览器和 Node.js 中运行。

## 项目简介

本项目展示了如何将 Go 代码编译为 WebAssembly，并在 JavaScript 环境中调用 Go 函数。主要功能包括：

- 使用标准 Go 编译器构建 WebAssembly
- 使用 TinyGo 构建体积更小的 WebAssembly
- 在浏览器和 Node.js 环境中测试 WebAssembly
- 提供基本的数学运算函数示例

## 环境准备

### 依赖项

- Go (1.11+)
- TinyGo
- Node.js 和 npm/yarn

### 安装 WebAssembly 支持文件

项目需要 WebAssembly 执行环境支持文件，可以通过以下命令获取：

```shell
# 标准 Go WebAssembly 执行环境
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./test/wasm_exec.js

# TinyGo WebAssembly 执行环境
cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" ./test/wasm_exec_tiny.js
```

### 安装项目依赖

```shell
npm install
# 或
yarn
```

## 构建

### 使用标准 Go 编译器构建

```shell
npm run build
# 或
make build
```

### 使用 TinyGo 编译器构建

```shell
npm run build:tiny
# 或
make build-tiny
```

## 测试

### 在浏览器中测试

#### 标准 Go WebAssembly

```shell
npm run test:browser
# 或
make serve
```

#### TinyGo WebAssembly

```shell
npm run test:browser:tiny
# 或
yarn run test:browser:tiny
```

### 在 Node.js 中测试

#### 标准 Go WebAssembly

> 注意：标准 Go WebAssembly 目前仅支持浏览器环境，不支持 Node.js

```shell
npm run test:node
```

#### TinyGo WebAssembly

```shell
npm run test:node:tiny
# 或
yarn run test:node:tiny
```

## 项目结构

```
.
├── main.go             # Go 源代码
├── test/               # 测试文件目录
│   ├── index.html      # 标准 Go WebAssembly 测试页面
│   ├── tiny.html       # TinyGo WebAssembly 测试页面
│   ├── tiny.cjs        # Node.js 测试脚本
│   ├── wasm_exec.js    # 标准 Go WebAssembly 执行环境
│   └── wasm_exec_tiny.js # TinyGo WebAssembly 执行环境
├── Makefile            # 构建脚本
└── package.json        # 项目配置
```

## 使用 Makefile

项目提供了 Makefile 以简化常用操作：

```shell
make build      # 构建标准 Go WebAssembly
make build-tiny # 构建 TinyGo WebAssembly
make serve      # 构建并启动 HTTP 服务器
make dev        # serve 的别名（通过依赖 serve 实现）
make start      # serve 的别名（通过显示执行 make serve）
make clean      # 清理构建文件
make deps       # 更新 Go 依赖
make test       # 运行测试
make help       # 显示帮助信息
```
