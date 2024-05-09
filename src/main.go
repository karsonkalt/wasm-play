package main

import (
    "syscall/js"
)

func main() {
    add := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        return args[0].Int() + args[1].Int()
    })

    js.Global().Set("add", add)

    select {}
}
