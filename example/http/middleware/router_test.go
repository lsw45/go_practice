package middleware

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"testing"
)

func TestRouter(t *testing.T) {
	h := http.NewServeMux()

	//从URL中过滤掉"/tmpfiles", 而剩下的路径是相对于根目录"/temp"的相对路径。
	h.Handle("/tempfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("temp"))))
	h.HandleFunc("/road", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("road"))
	})
	startAdmin(h)

	chain := NewChain(
		chain1(),
	)
	mw := chain.Then(h)
	http.Handle("/", mw)

	l, _ := net.Listen("tcp", "9900")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt) // kill -2 <pid>
	<-sig
	_ = l.Close()
}

func startAdmin(h *http.ServeMux) {
	h.Handle("/marketing/", AdminRoute())
}
