package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHead(t *testing.T) {
	//获取get的请求参数
	var r = http.Request{}
	_ = r.ParseForm()
	if len(r.Form) <= 0 || len(r.Form["merchantCode"]) <= 0 {
		return
	}
	merchantCode := r.Form["merchantCode"][0]
	//或者直接
	merchantCode = r.FormValue("merchantCode")

	//获取post表单
	if len(r.PostForm) <= 0 || len(r.PostForm["merchantCode"]) <= 0 {
		return
	}
	merchantCode = r.PostForm["merchantCode"][0]
	//或者直接
	merchantCode = r.PostFormValue("merchantCode")

	fmt.Println(merchantCode)
}
