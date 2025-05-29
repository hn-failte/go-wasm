package main

import "syscall/js"

func goAdd(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		js.Global().Get("Error").New("Invalid number of arguments. Expected 2.").Call("throw")
		return nil
	}

	num1 := args[0].Float()
	num2 := args[1].Float()
	result := num1 + num2
	return result
}

func goMinus(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		js.Global().Get("Error").New("Invalid number of arguments. Expected 2.").Call("throw")
		return nil
	}

	num1 := args[0].Float()
	num2 := args[1].Float()
	result := num1 - num2
	return result
}

func goMultiply(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		js.Global().Get("Error").New("Invalid number of arguments. Expected 2.").Call("throw")
		return nil
	}

	num1 := args[0].Float()
	num2 := args[1].Float()
	result := num1 * num2
	return result
}

func goDiv(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		js.Global().Get("Error").New("Invalid number of arguments. Expected 2.").Call("throw")
		return nil
	}

	num1 := args[0].Float()
	num2 := args[1].Float()
	result := num1 / num2
	return result
}

func goIntDiv(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		js.Global().Get("Error").New("Invalid number of arguments. Expected 2.").Call("throw")
		return nil
	}

	num1 := args[0].Int()
	num2 := args[1].Int()
	result := num1 / num2
	return result
}

func goMod(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		js.Global().Get("Error").New("Invalid number of arguments. Expected 2.").Call("throw")
		return nil
	}

	num1 := args[0].Int()
	num2 := args[1].Int()
	result := num1 % num2
	return result
}

func gogogo(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		js.Global().Get("Error").New("Invalid number of arguments. Expected 2.").Call("throw")
		return nil
	}

	num1 := args[0].Int()
	num2 := args[1].Int()
	result := 10 * num1 * num2
	return result
}

func main() {
	// 挂载方法到全局
	js.Global().Set("gogogo", js.FuncOf(gogogo))

	// 设置对象的属性或方法
	obj := js.Global().Get("Object").New()
	obj.Set("add", js.FuncOf(goAdd))
	obj.Set("minus", js.FuncOf(goMinus))
	obj.Set("multiply", js.FuncOf(goMultiply))
	obj.Set("div", js.FuncOf(goDiv))
	obj.Set("intDiv", js.FuncOf(goIntDiv))
	obj.Set("mod", js.FuncOf(goMod))
	// 挂载对象到全局
	js.Global().Set("go", obj)
	done := make(chan struct{}, 0)
	<-done
}
