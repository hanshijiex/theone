package main

import (
	"fmt"
	"theone/engine"
)

type worker struct {

}

func (worker *worker) Run(req *engine.Request, res *engine.Response) error {
	readBuf := req.GetReadBuf()
	len := req.GetReadLen()
	res.SetWriteBuf(readBuf)
	res.SetWriteLen(len)
	return nil
}

func main() {
	fmt.Println("hello world")
	s := engine.NewServer()
	s.SetController(&worker{})
	s.Serve()
}
