package main

import "syscall/js"

func goAdd(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "Invalid number of arguments. Expected 2."
	}

	num1 := args[0].Float()
	num2 := args[1].Float()
	result := num1 + num2
	return result
}

func goMinus(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "Invalid number of arguments. Expected 2."
	}

	num1 := args[0].Float()
	num2 := args[1].Float()
	result := num1 - num2
	return result
}

func goMultiply(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "Invalid number of arguments. Expected 2."
	}

	num1 := args[0].Float()
	num2 := args[1].Float()
	result := num1 * num2
	return result
}

func goDiv(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "Invalid number of arguments. Expected 2."
	}

	num1 := args[0].Float()
	num2 := args[1].Float()
	result := num1 / num2
	return result
}

func goMod(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "Invalid number of arguments. Expected 2."
	}

	num1 := args[0].Int()
	num2 := args[1].Int()
	result := num1 % num2
	return result
}

func gogogo(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "Invalid number of arguments. Expected 2."
	}

	num1 := args[0].Int()
	num2 := args[1].Int()
	result := 10 * num1 * num2
	return result
}

func main() {
	js.Global().Set("goAdd", js.FuncOf(goAdd))
	js.Global().Set("goMinus", js.FuncOf(goMinus))
	js.Global().Set("goMultiply", js.FuncOf(goMultiply))
	js.Global().Set("goDiv", js.FuncOf(goDiv))
	js.Global().Set("goMod", js.FuncOf(goMod))
	js.Global().Set("gogogo", js.FuncOf(gogogo))
	done := make(chan struct{}, 0)
	<-done
}
