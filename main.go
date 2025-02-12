package main

import (
	"fmt"
	"goproject/context"
	"goproject/server"
	"runtime"
)

func main() {

	fmt.Printf("Go version: %s\n", runtime.Version())

	context.InitContext()

	server.Serve()
}
