package test

import (
	"fmt"
	"net"
	"testing"
	"theone/server"
	"time"
)

func TestServer(t *testing.T) {
	go ClientTest()
	s := server.NewServer()
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
		buf := make([]byte, 1024)
		lens, err := conn.Read(buf)
		if err != nil {
			fmt.Print("read buf err:", err)
		}
		fmt.Println("recv data: ", string(buf[:lens]))
		time.Sleep(1*time.Second)
	}

}