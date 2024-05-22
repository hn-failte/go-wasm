const fs = require('fs')
const path = require('path')
require('./wasm_exec_tiny.js')

const go = new global.Go()

const wasmBuffer = fs.readFileSync(path.resolve(__dirname, './main_tiny.wasm'))

WebAssembly.instantiate(
  wasmBuffer,
  go.importObject
).then(res => {
  go.run(res.instance)
  return res.instance
}).then(res => {
  console.log('res', res)
})
