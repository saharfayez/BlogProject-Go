package main

import (
	"fmt"
	_ "goproject/internal/app/context/impl"
	"goproject/internal/app/server"
	"runtime"
)

func main() {

	fmt.Printf("Go version: %s\n", runtime.Version())

	server.Serve()
}
