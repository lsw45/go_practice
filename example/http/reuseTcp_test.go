package http

import (
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestReuseTcp(t *testing.T) {
	count := 100
	for i := 0; i < count; i++ {
		resp, err := http.Get("https://www.oschina.net") //这个站点用的是 HTTPS，所以重用了 TCP 的话，那么一次建立 TLS 连接后面就不用重建了，非常方便观察。
		if err != nil {
			panic(err)
		}

		io.Copy(ioutil.Discard, resp.Body) //必须将http.Response的Body读取完毕并且关闭后，才会重用TCP连接
		resp.Body.Close()
	}
}
