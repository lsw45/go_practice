package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	// "bufio"
	// "io"
	// "log"
	"os"
	"os/signal"
	"sort"
)

func TestXX(t *testing.T) {

	sortStruct()
	c := make(chan os.Signal)
	signal.Notify(c)
	//signal.Notify(c, syscall.SIGHUP, syscall.SIGUSR2)  //监听指定信号
	s := <-c //阻塞直至有信号传入
	fmt.Println("get signal:", s)

	// 一次性读文本全部内容
	/*
		f, err := os.Open("appConfig.txt")
		if err != nil {
			fmt.Printf("%q", err)
		}
		defer f.Close()

		// var b bytes.Buffer
		all, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Printf("%q", err)
		}
		fmt.Println(string(all))
	*/

	/*
		f, _ := ioutil.ReadFile("appConfig.txt")
		fmt.Println(string(f))
	*/

	/*
		var buf bytes.Buffer
		f, _ := os.Open("appConfig.txt")
		defer f.Close()
		all, _ := ioutil.ReadAll(f)
		buf.Write(all)
		fmt.Println(buf.String())
	*/

	// 循环读取文件
	/*
		f, err := os.Open("appConfig.txt")
		if err != nil {
			fmt.Printf("%q", err)
		}
		defer f.Close()

		buf := make([]byte, 1000)
		bfrd := bufio.NewReader(f)
		count := 0 //统计循环读取次数
		for {
			n, err := bfrd.Read(buf)
			os.Stdout.Write(buf[:n])
			count++
			if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
				if err == io.EOF {
					fmt.Println("io.EOF；", count)
					os.Exit(2)
				}
				fmt.Printf("%q", err)
			}
		}
	*/
}

type addr struct {
	Province string
	City     string
}
type stu struct {
	Name string
	Age  int
	Addr addr
}

// 首字母排序
func sortStruct() {
	js := `{"Age":18,"name":"xiaoming","Addr":{"Province":"Hunan","City":"ChangSha"}}` //name是小写
	// 1、转化为结构体
	var stu1 stu
	json.Unmarshal([]byte(js), &stu1)

	// 2、读取为字节
	b, _ := json.Marshal(stu1)
	fmt.Println(string(b))

	// 转化为map
	var m map[string]interface{}
	// json.Unmarshal(b, &m)
	json.Unmarshal([]byte(js), &m)
	fmt.Println(m)

	// 排序
	newMp := make([]string, 0)
	for k, _ := range m {
		newMp = append(newMp, k)
	}
	sort.Strings(newMp)
	for _, v := range newMp {
		s, _ := json.Marshal(m[v])
		fmt.Println(v, ":", string(s))
	}
}
