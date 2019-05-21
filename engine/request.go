package engine

import (
	"fmt"
	"io"
)

type Request struct {
	readBuf []byte
	readLen int
}

func (request *Request) ReadFrom(r io.Reader) error {
	len, err := r.Read(request.readBuf)
	if err != nil {
		fmt.Println("read from conn err: ", err)
		return err
	}
	fmt.Println("read len: ", len)
	request.readLen = len
	return nil
}

func (request *Request) GetReadBuf() []byte {
	return request.readBuf
}

func (request *Request) GetReadLen() int {
	return request.readLen
}