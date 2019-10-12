package http

import (
	"fmt"
	_ "io/ioutil"
	_ "net/http"
	"net/url"
	_ "strings"
	"testing"
)

func TestJoin(t *testing.T) {
	// athEscape escapes the string so it can be safely placed inside a URL path segment.
	fmt.Println(url.QueryEscape("我爱你~~~"))

	v := url.Values{}
	v.Add("msg", "此订单不存在或已经提ssssss  xxxx")
	body := v.Encode()
	fmt.Println(v)
	fmt.Println(body)
	// url decode
	m, _ := url.ParseQuery(body)
	fmt.Println(m)

	i := "xxt三毛"

	fmt.Println(i)
	fmt.Printf("%x\n", []rune(i))
	fmt.Printf("%x\n", []byte(i))
}
