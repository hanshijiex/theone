package server

import (
	"fmt"
	"net"
)

type Server struct {
	// 服务器名称
	Name string
	// IP版本
	IpVersion string
	// ip地址
	Ip string
	// 端口号
	Port uint16
	// 状态 关闭-1 运行0
	Status int8
}

func NewServer() *Server {
	s := &Server{
		Name:      "Serve",
		IpVersion: "tcp4",
		Ip:        "127.0.0.1",
		Port:      8082,
		Status:    -1,
	}
	return s
}

func (server *Server) Start() {
	fmt.Printf("[Server %s] Starting...\n", server.Name)
	fmt.Println(fmt.Sprintf("%s:%d", server.Ip, server.Port))
	addr, err := net.ResolveTCPAddr(server.IpVersion, fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("resolve tcp addr err: ", err)
		return
	}
	listener, err := net.ListenTCP(server.IpVersion, addr)
	if err != nil {
		fmt.Println("listen to addr err: ", addr.String)
	}
	server.Status = 0
	fmt.Printf("[TheOne Server Started ip:%s port:%d]\n", server.Ip, server.Port)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("accept ")
		}
		go func() {
			buf := make([]byte, 1024)
			lens, err := conn.Read(buf)
			if err != nil {
				fmt.Println("read from conn err:", err)
			}
			fmt.Printf("recv data:%s\n", string(buf[:lens]))
			if _, err := conn.Write(buf[:lens]); err != nil {
				fmt.Println("conn Write err: ", err)
			}
		}()
	}
}

func (server *Server) Stop() {
	server.Status = -1
	fmt.Printf("[Server %s\n] Stoped", server.Name)
}

func (server *Server) Serve() {
	server.Start()
	select {}
}
