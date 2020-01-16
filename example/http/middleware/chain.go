package middleware

import (
	"net/http"
)

type Adapter func(handler http.Handler) http.Handler

type Chain struct {
	adapter []Adapter
}

func NewChain(adapters ...Adapter) Chain {
	return Chain{adapter: append(([]Adapter)(nil), adapters...)}
}

func (c Chain) Then(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}

	for i := range c.adapter {
		h = c.adapter[len(c.adapter)-1-i](h)
	}
	return h
}

func chain1() Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					_, _ = w.Write([]byte("error"))
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
