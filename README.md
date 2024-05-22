# go-wasm

A demo project of golang wasm.

## scripts

### build

build normal go wasm
```shell
npm run build
```

build normal tinygo wasm
```shell
npm run build:tiny
```

### test

test go wasm in browser
```shell
npm run test:browser
```

test tinygo wasm in browser
```shell
yarn run test:browser:tiny
```

test go wasm in node. But go wasm now only support browser exclude node.
```shell
npm run test:node
```

test tinygo wasm in node
```shell
yarn run test:node:tiny
```

## env

Here is the source of “wasm_exec.js” file.

```shell
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./wasm_exec.js
```

Here is the source of “wasm_exec_tiny.js” file.

```shell
cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" ./wasm_exec_tiny.js
```
