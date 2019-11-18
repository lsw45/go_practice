package main

import (
	"log"
	"net/http"
)

type httpServer struct {
}

func (server httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("xxxxx"))
	w.Write([]byte(r.URL.Path))
}

func main() {
	var server httpServer
	http.Handle("/", server)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}
