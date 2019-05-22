package engine

import (
	"fmt"
	"io"
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
	// 控制器
	Controller IController
}

func NewServer() *Server {
	s := &Server {
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
		originalConn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("accept ")
		}

		conn := &Conn {
			OriginalConn: originalConn,
			Request: &Request {
				readBuf: make([]byte, 1024),
				readLen: 0,
			},
			Response: &Response {
				writeBuf: make([]byte, 1024),
				writeLen: 0,
			},
			Status: 1,
		}

		go func() {
			for {
				err := conn.Request.ReadFrom(conn.OriginalConn)
				if err == io.EOF {
					fmt.Println("conn is closed")
					conn.OriginalConn.Close()
					conn.Status = 0
					break
				}

				err = server.Controller.Run(conn.Request, conn.Response)
				if err != nil {
					fmt.Println("call controller run err: ", err)
				}

				if conn.isActive() {
					err = conn.Response.WriteTo(conn.OriginalConn)
				}
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

func (server *Server) SetController(controller IController) {
	server.Controller = controller
}
