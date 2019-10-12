package main

import (
	//"helloworld/hello"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

const dataFile = "data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}
type Message struct {
	Name, Text string
}

func TestDecodeJson(t *testing.T) {
	log.Println("————————————main——————————————")
	os_File_Json()
	os_File_Json1()
	byteUnmarshal()
	streamDecoder()
	fileStream()
	JsonStream()
}

// stream，比如文件流，http request。
// 如下：从流里把数据按Message结构一个一个读出来
func streamDecoder() {
	log.Println("————————————streamDecoder——————————————")
	const jsonStream = `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]
`

	dec := json.NewDecoder(strings.NewReader(jsonStream)) // 流：bytes.NewReader([]byte(jsonStream))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	// t的类型是json.Delim
	fmt.Printf("start:%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.Name, m.Text)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	// t的类型是json.Delim
	fmt.Printf("%T: %v:end\n", t, t)
}

// 数组元素的地址:[]*Feed
func os_File_Json() {
	log.Println("————————————os_File_Json——————————————")
	file, err := os.Open(dataFile)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	fmt.Printf("%T | %v\n", feeds, feeds) //[]*main.Feed | [0x12ebd940 0x12ebd9c0 0x12ebd9e0]

	for i := range feeds {
		fmt.Println(*feeds[i])
	}
}

// 数组元素的拷贝:[]Feed
func os_File_Json1() {
	log.Println("————————————os_File_Json1——————————————")
	file, err := os.Open(dataFile)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var feeds []Feed
	err = json.NewDecoder(file).Decode(&feeds)

	fmt.Printf("%T | %v\n", feeds, feeds) //[]main.Feed | [{npr http://www.npr.org/rss/rss.php?id=1001 rss} {npr http://www.npr.org/rss/rss.php?id=1008 rss} {npr http://www.npr.org/rss/rss.php?id=1006 rss}]

	for i := range feeds {
		fmt.Println(feeds[i])
	}
}

// 文件流
func fileStream() {
	log.Println("————————————fileStream——————————————")
	file, err := os.Open(dataFile)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file) //文件流
	dec := json.NewDecoder(reader)

	// 知识点：Token包括：json.Delim，基本类型（bool，float64，Number，string）和nil。其中json.Delim包括[ ] { }
	// read open bracket读开括号
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	// t的类型是json.Delim
	fmt.Printf("start:%T: %v\n", t, t) //start:json.Delim: [

	// while the array contains values
	for dec.More() {
		var m Feed
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.Name, m.URI)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	// t的类型是json.Delim
	fmt.Printf("%T: %v:end\n", t, t)

	var feeds []Feed
	err = json.NewDecoder(file).Decode(&feeds)

	fmt.Printf("%T | %v\n", feeds, feeds) //[]main.Feed | [{npr http://www.npr.org/rss/rss.php?id=1001 rss} {npr http://www.npr.org/rss/rss.php?id=1008 rss} {npr http://www.npr.org/rss/rss.php?id=1006 rss}]

	for i := range feeds {
		fmt.Println(feeds[i])
	}
}

type Animal struct {
	Name  string
	Order string
}

// 以[]byte存在于内存中的用json.Unmarshal
func byteUnmarshal() {
	log.Println("————————————byteUnmarshal——————————————")

	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)
}

// 下面我们将使用这两个结构体来演示自定义类型的编码和解码。
type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func JsonStream() {
	// r 被定义为一个空接口。json.Unmarshal() 函数将一个 JSON 对象解码到空接口 r 中，最终 r 将会是一个键值对的 map[string]interface{} 结构
	log.Println("————————————JsonStream——————————————")

	var r interface{}
	b := []byte(`{
	    "Title": "Go语言编程",
	    "Authors": ["XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan",
	    "XuDaoli"],
	    "Publisher": "ituring.com.cn",
	    "IsPublished": true,
	    "Price": 9.99,
	    "Sales": 1000000
	}`)
	json.Unmarshal(b, &r)
	fmt.Println(r)
	// os.Exit(2)

	// 首先我们来看一下基本数据类型到 JSON 字符串的编码过程。
	// 这里是一些原子值的例子。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// 这里是一些切片和 map 编码成 JSON 数组和对象的例子。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON 包可以自动的编码你的自定义类型。
	// 编码仅输出可导出的字段，并且默认使用他们的名字作为 JSON 数据的键。
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 你可以给结构字段声明标签来自定义编码的 JSON 数据键名称。
	// 在上面 Response2 的定义可以作为这个标签这个的一个例子。
	res2D := Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// 现在来看看解码 JSON 数据为 Go 值的过程。
	// 这里是一个普通数据结构的解码例子。
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// 我们需要提供一个 JSON 包可以存放解码数据的变量。
	// 这里的 map[string]interface{} 将保存一个 string 为键，值为任意值的map。
	var dat map[string]interface{}

	// 这里就是实际的解码和相关的错误检查。
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 为了使用解码 map 中的值，我们需要将他们进行适当的类型转换。
	// 例如这里我们将 num 的值转换成 float64类型。
	num := dat["num"].(float64)
	fmt.Println(num)

	// 访问嵌套的值需要一系列的转化。
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 我们也可以解码 JSON 值到自定义类型。
	// 这个功能的好处就是可以为我们的程序带来额外的类型安全加强，并且消除在访问数据时的类型断言。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 在上面的例子中，我们经常使用 byte 和 string 作为使用标准输出时数据和 JSON 表示之间的中间值。
	// 我们也可以和os.Stdout 一样，直接将 JSON 编码直接输出至 os.Writer流中，或者作为 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	fmt.Println("xx")
	enc.Encode(d)
}
