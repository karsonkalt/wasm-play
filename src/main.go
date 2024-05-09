package main

import (
    "syscall/js"
)

var counter int

func main() {
    add := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        return args[0].Int() + args[1].Int()
    })

    js.Global().Set("add", add)

    multiply := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        return args[0].Int() * args[1].Int()
    })

    js.Global().Set("multiply", multiply)

    divide := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        return args[0].Int() / args[1].Int()
    })

    js.Global().Set("divide", divide)

    subtract := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        return args[0].Int() - args[1].Int()
    })

    js.Global().Set("subtract", subtract)

    increment := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        counter++
        return nil
    })

    js.Global().Set("increment", increment)

    getCounter := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        return counter
    })

    js.Global().Set("getCounter", getCounter)

    printMessage := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        go func() {
            for {
                fmt.Println("Hello from Go!")
                time.Sleep(time.Second)
            }
        }()
        return nil
    })

    js.Global().Set("printMessage", printMessage)

    select {}
}
