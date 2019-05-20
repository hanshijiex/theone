package engine

import "io"

type request struct {
	readBuf []byte
}

func (request *request) ReadFrom(r io.Reader) {

}