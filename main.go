package main

import (
	"fmt"
	"theone/engine"
)

func main() {
	fmt.Println("hello world")
	s := engine.NewServer()
	s.Serve()
}
