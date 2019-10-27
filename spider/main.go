package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main(){
	resp,err := http.Get("http://www.baidu.com")
	if err !=nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("Error:status code:",resp.StatusCode)
	}


	//gbk编码改utf ：utf8Reader :=transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
	e := determineEncoding(resp.Body)
	utf8Reader :=transform.NewReader(resp.Body,e.NewDecoder())
	all,err :=ioutil.ReadAll(utf8Reader)
	if err !=nil {
		panic(err)
	}
	fmt.Printf("%s",all)
}

//判断网站是什么编码:GBK,UTF8……
func determineEncoding(r io.Reader)encoding.Encoding{
	// resp.Body只能被解析一次，所以这里用peek，截取
	b,err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	// DetermineEncoding determines the encoding of an HTML document by examining up to the first 1024 bytes of content and the declared Content-Type.
	e,_,_:=charset.DetermineEncoding(b,"")
	return e
}