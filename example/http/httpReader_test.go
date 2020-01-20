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
	values, err := url.ParseRequestURI("https://www.baidu.com/s?wd=%E6%90%9C%E7%B4%A2&rsv_spt=1&issp=1&f=8&rsv_bp=0&rsv_idx=2")
	fmt.Println(values) // https://www.baidu.com/s?wd=%E6%90%9C%E7%B4%A2&rsv_spt=1&issp=1&f=8&rsv_bp=0&rsv_idx=2

	if err != nil {
		fmt.Println(err)
	}

	urlParam := values.RawQuery // url的 ? 之后的部分
	fmt.Println(urlParam)       // wd=%E6%90%9C%E7%B4%A2&rsv_spt=1&issp=1&f=8&rsv_bp=0&rsv_idx=2

	// ParseQuery传入的必须是参数，也就是url里边的RawQuery的值 就是url?之后的path
	fmt.Println(url.ParseQuery(urlParam)) //map[f:[8] issp:[1] rsv_bp:[0] rsv_idx:[2] rsv_spt:[1] wd:[搜索]] <nil>

	//url.Query()直接就解析成map了，因为内部代码就是封装上面的两步
	urlValue := values.Query() // 和下面的c变量类型相同都为url.Values类型，有相同的属性方法
	fmt.Println(urlValue)      // map[f:[8] issp:[1] rsv_bp:[0] rsv_idx:[2] rsv_spt:[1] wd:[搜索]]

	//val := url.Values{}
	c := url.Values{"method": {"get", "put"}, "id": {"1"}}
	fmt.Println(c) // map[id:[1] method:[get put]]

	// encodes the values into ``URL encoded'' form ("bar=baz&foo=quux") sorted by key.
	fmt.Println(c.Encode())      // id=1&method=get&method=put
	fmt.Println(c.Get("method")) // get-只能获取到第一个元素

	c.Set("method", "post") // 修改method的值为post
	fmt.Println(c)          //map[id:[1] method:[post]]

	c.Del("method")    // 删除method元素
	c.Add("new", "hi") // 添加新的元素new:hi
	fmt.Println(c)     //map[id:[1] new:[hi]]

	// QueryEscape escapes the string so it can be safely placed inside a URL query.
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
}
