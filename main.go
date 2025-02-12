package main

import (
	"fmt"
	"goproject/server"
	"runtime"
)

func main() {

	fmt.Printf("Go version: %s\n", runtime.Version())

	server.Serve()
}
