package main

import "fmt"
import "net/url"
import "strings"

func main() {

	// urlParse()
	urlEncode()
	// 我们将解析这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
	s := "https://memtest.ipay.so/lingxi/cardMemActivity/payBack?bankType%3DCFT%26busicd%3DWPAY%26channelOrderNum%3D4200000450201910256752010308%26charset%3Dutf-8%26chcd%3DWXP%26chcdDiscount%3D0.00%26consumerAccount%3DorS1BuPK8d1avXUCNVrEgf2asPJo%26errorDetail%3DSUCCESS%26inscd%3D10130001%26mchntid%3D013361641120005%26merDiscount%3D0.00%26orderNum%3Drecharge_19102520272923961%26payTime%3D2019-10-25+20%3A27%3A35%26respcd%3D00%26signType%3DSHA256%26terminalid%3D87602350049134%26transTime%3D2019-10-25T20%3A27%3A29%2B08%3A00%26txamt%3D000000000001%26txndir%3DA%26version%3D2.2%26sign%3De5e6c68a658ccde13acff49c0f315097ac53c01320572ebe49fb27b16dedb6b6"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme) //postgres

	raw, err := url.QueryUnescape(u.RawQuery)
	fmt.Println(raw) //k=v
	value, err := url.ParseQuery(raw)
	fmt.Println(value) //map[k:[v]]
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
	s := "https://note.youdao.com/web?a=11&b=22&c=33#/file/recent/note/1f41f1a02252a3258e70fef5d6ade7ed/#ttl"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)                   // https
	fmt.Println(u.Host)                     // note.youdao.com
	fmt.Println(u.Path)                     // /web
	fmt.Println(u.Fragment)                 // /file/recent/note/1f41f1a02252a3258e70fef5d6ade7ed/#ttl
	fmt.Println(u.RawQuery)                 // a=11&b=22&c=33
	fmt.Println(url.ParseQuery(u.RawQuery)) // map[a:[11] b:[22] c:[33]] <nil>
}

func urlEncode() {
	u := url.Values{}

	sso := struct {
		Next  string
		AppID string
	}{
		Next:  "/lingxi/platform/auth?",
		AppID: "a.Base.App.Config().Sso.AppId",
	}

	u.Add("orderNum", "recharge_19102520272923961")
	u.Add("next", sso.Next)
	fmt.Printf("%s\n", u.Encode())
}
