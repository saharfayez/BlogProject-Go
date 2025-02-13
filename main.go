package main

import (
	"fmt"
	_ "goproject/impl/context"
	"goproject/server"
	"runtime"
)

func main() {

	fmt.Printf("Go version: %s\n", runtime.Version())

	server.Serve()
}
