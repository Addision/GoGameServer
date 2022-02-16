package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"os"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	// if err := L.DoString(`print("hello world")`); err != nil {
	// 	panic(err)
	// }

	luaPath, _ := os.Getwd()
	luaPath = luaPath + "/src/gotest/fib.lua"
	fmt.Println(luaPath)
	if err := L.DoFile(luaPath); err != nil {
		panic(err)
	}

	err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("fib"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(10))
	if err != nil {
		panic(err)
	}

	ret := L.Get(-1)
	L.Pop(1)
	res, ok := ret.(lua.LNumber)
	if ok {
		fmt.Println(int(res))
	} else {
		fmt.Println("unexpected result")
	}
}
