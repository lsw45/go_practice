package main

import "fmt"
import "net/url"
import "strings"

func main() {

	urlParse()
	// 我们将解析这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	return

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())

	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}

func urlParse() {
	s := "https://note.youdao.com/web/#/file/recent/note/1f41f1a02252a3258e70fef5d6ade7ed/"
	// s := "https://note.youdao.com/web/file/recent/note/1f41f1a02252a3258e70fef5d6ade7ed/"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)                   // https
	fmt.Println(u.Host)                     // note.youdao.com
	fmt.Println(u.Path)                     // /web/
	fmt.Println(u.Fragment)                 // /file/recent/note/1f41f1a02252a3258e70fef5d6ade7ed/
	fmt.Println(url.ParseQuery(u.RawQuery)) // map[] <nil>
}
