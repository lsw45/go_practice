package pool

import (
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

type ConnRes interface {
	Close() error
}
type Factory func() (ConnRes, error)

type Pool struct {
	conns   chan ConnRes
	factory Factory
}

func GetHttpPool() {
	//con1, _ := DefaultHttpPool.newConn()
	route := mux.NewRouter()
	http.Handle("/lingxi", route)
	server := &http.Server{Addr: "8080"}
	_ = server.ListenAndServe()

	con1 := DefaultHttpPool.Get()
	_ = con1.Close()
}

var DefaultHttpPool = NewPool(func() (ConnRes, error) {
	return net.Dial("tcp", ":8080")
}, 10)

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
		conn, _ = p.newConn()
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
