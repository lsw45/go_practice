package middleware

import "net/http"

// Route 后台管理的请求统一入口
func AdminRoute() (mux *http.ServeMux) {
	mux = http.NewServeMux()

	mux.HandleFunc("/admin/route", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("admin"))
	})

	return mux
}
