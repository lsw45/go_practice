package Atomic

import (
	"fmt"
	"net"
	"sync/atomic"
	"testing"
)

func TestLoad32(t *testing.T) {
	s := int32(1)
	a := atomic.LoadInt32(&s) != 0
	t.Log(a)
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	t.Log(err)

	listner := ln.(*net.TCPListener)
	t.Logf("%+v", listner)
	fmt.Println(222)
	conn, err := listner.Accept()
	fmt.Println(333)
	t.Log(err)
	t.Logf("%+v", conn)
}
