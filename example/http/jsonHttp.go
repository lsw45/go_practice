package http

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "bufio"
	// "io"
	"github.com/astaxie/beego/httplib"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestXX(t *testing.T) {
	httpClient()
	jsonDecoder()
	closeHttpBody()
}

//
func httpClient() {
	b := httplib.Post("http://beego.me/")
	b.Param("username", "astaxie")
	b.Param("password", "123456")
	b.PostFile("uploadfile2", "httplib.txt")
	str, err := b.String()
	if err != nil {
		log.Fatal(err)
	}

	fileName := "httpClient.html"

	var f *os.File
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		f, _ = os.Create(fileName) //创建文件
	} else {
		//打开文件并清空
		f, _ = os.Open(fileName)
		f.WriteString("")
	}

	defer f.Close()
	f.WriteString(str)
	// fmt.Println(str)
}

func jsonDecoder() {
	resp, _ := http.Get("http://www.baidu.com")
	if resp == nil {
		log.Fatal("empty response")
	}
	defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	// os.Exit(2)

	decoder := json.NewDecoder(resp.Body)
	var obj interface{}
	if err := decoder.Decode(&obj); err != nil {
		fmt.Println("请求体反序列化失败：", err)
	}
	fmt.Printf("%T\n", &obj)
	fmt.Println(obj)
}

func closeHttpBody() {
	resp, err := http.Get("https://api.ipify.orformat=json")
	// defer resp.Body.Close() //not ok
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
	}

	fmt.Println("body:", string(body))
}
