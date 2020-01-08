package http

import (
	"fmt"
	_ "io/ioutil"
	"log"
	"net/http"
	_ "net/http"
	"net/url"
	_ "strings"
	"testing"
)

func TestJoin(t *testing.T) {
	// athEscape escapes the string so it can be safely placed inside a URL path segment.
	fmt.Println(url.QueryEscape("我爱你~~~")) //%E6%88%91%E7%88%B1%E4%BD%A0~~~

	v := url.Values{}
	v.Add("msg", "此订单不存在或已经提ssssss  xxxx")
	body := v.Encode()
	fmt.Println(v)    //map[msg:[此订单不存在或已经提ssssss  xxxx]]
	fmt.Println(body) //msg=%E6%AD%A4%E8%AE%A2%E5%8D%95%E4%B8%8D%E5%AD%98%E5%9C%A8%E6%88%96%E5%B7%B2%E7%BB%8F%E6%8F%90ssssss++xxxx
	// url decode
	m, _ := url.ParseQuery(body)
	fmt.Println(m) //map[msg:[此订单不存在或已经提ssssss  xxxx]]

	i := "xxt三毛"
	fmt.Println(i)
	fmt.Printf("%x\n", []rune(i)) //[78 78 74 4e09 6bdb]
	fmt.Printf("%x\n", []byte(i)) //787874e4b889e6af9b

	var w http.ResponseWriter
	//http.ResponseWriter 是否实现了 http.Pusher 接口来判断是否支持 Server Push。
	if push, ok := w.(http.Pusher); ok {
		if err := push.Push("/app.js", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
	w.Header().Add()
}
