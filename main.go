package main

import (
	"fmt"
	"theone/server"
)

func main() {
	fmt.Println("hello world")
	s := server.NewServer()
	s.Serve()
}
