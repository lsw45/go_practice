package pool

import (
	"fmt"
	"net"
)

type ConnRes interface {
	Close() error
}
type Factory func() (ConnRes, error)

type Pool struct {
	conns   chan ConnRes
	factory Factory
}

func ExampleGetHttpPool() {

}

var DefaultHttpPool = NewPool(func() (ConnRes, error) {
	return net.Dial("tcp", ":8080")
}, 10)

func init() {
	for i := 0; i < 10; i++ {
		conn, _ := DefaultHttpPool.newConn()
		DefaultHttpPool.Put(conn)
	}
}

func NewPool(factory Factory, cap int) *Pool {
	return &Pool{
		conns:   make(chan ConnRes, cap),
		factory: factory,
	}
}

func (p *Pool) newConn() (ConnRes, error) {
	return p.factory()
}

func (p *Pool) Get() (conn ConnRes) {
	select {
	case conn = <-p.conns:
		{
		}
	default:
		fmt.Println("the pool of conn is empty")
		conn, _ = p.newConn()
		p.conns <- conn
	}
	return
}

func (p *Pool) Put(conn ConnRes) {
	select {
	case p.conns <- conn:
		{
		}
	default:
		conn.Close()
	}
}
