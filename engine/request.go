package engine

import (
	"fmt"
	"io"
)

type request struct {
	readBuf []byte
	readLen int
}

func (request *request) ReadFrom(r io.Reader) error {
	len, err := r.Read(request.readBuf)
	if err != nil {
		fmt.Println("read from conn err: ", err)
	}
	fmt.Println("read len: ", len)
	request.readLen = len
}