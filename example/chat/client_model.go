package main

import "net"

type Client struct {
	//客户端连接
	conn net.Conn
	//昵称
	name string
	//远程地址
	addr string
}
