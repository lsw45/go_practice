package main

import (
	"net/http"
)

func dirRoute(w http.ResponseWriter, r *http.Request) {
	// 参数必须是[]byte
	w.Write([]byte(http.Dir(".")))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/dir", dirRoute)
	http.ListenAndServe(":8080", nil)
}
