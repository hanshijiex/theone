package engine

import (
	"fmt"
	"io"
)

type Response struct {
	writeBuf []byte
	writeLen int
}

func (response *Response) WriteTo(w io.Writer) error {
	len, err := w.Write(response.writeBuf[:response.writeLen])
	if err != nil {
		fmt.Println("write to conn err: ", err)
		return err
	}
	fmt.Println("read len: ", len)
	return nil
}

func (response *Response) SetWriteBuf(writeBuf []byte) {
	response.writeBuf = writeBuf
}

func (response *Response) SetWriteLen(len int) {
	response.writeLen = len
}