package engine

import "net"

type Conn struct {
	Id uint32
	// 请求
	Request *Request
	// 响应
	Response *Response
	// 状态 close0 open1
	Status int8
	// 原始连接
	OriginalConn *net.TCPConn
}

func (conn *Conn) isAlive() bool {
	return conn.Status != 0
}