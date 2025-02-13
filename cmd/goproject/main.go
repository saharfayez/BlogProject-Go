package main

import (
	"fmt"
	_ "goproject/internal/app/impl/context"
	"goproject/internal/app/server"
	"runtime"
)

func main() {

	fmt.Printf("Go version: %s\n", runtime.Version())

	server.Serve()
}
