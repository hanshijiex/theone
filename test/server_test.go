package test

import (
	"fmt"
	"net"
	"testing"
	"theone/engine"
	"time"
)

func TestServer(t *testing.T) {
	go ClientTest()
	s := engine.NewServer()
	s.Serve()
}

func ClientTest() {
	conn, err := net.Dial("tcp4", "127.0.0.1:8082")
	if err != nil {
		fmt.Println("Dial err: ", err)
	}
	for {
		_, err := conn.Write([]byte("hello world"))
		if err != nil {
			fmt.Println("write err: ", err)
		}
		fmt.Println("client write hello world ")
		buf := make([]byte, 1024)
		lens, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err:", err)
		}
		fmt.Println("client recv data: ", string(buf[:lens]))
		time.Sleep(1*time.Second)
	}

}
