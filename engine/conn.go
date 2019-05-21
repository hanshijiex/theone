package engine

import "net"

type Conn struct {
	// 请求
	Request *request
	// 响应
	Response *response
	// 控制器
	Controller *controller
	// 状态
	Status int8
	// 原始连接
	OriginalConn *net.TCPConn
}