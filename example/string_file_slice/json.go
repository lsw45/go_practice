package main

import "encoding/json"
import "fmt"
import "os"

// 下面我们将使用这两个结构体来演示自定义类型的编码和解码。
type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	b := []byte(`{
	    "Title": "Go语言编程",
	    "Authors": ["XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan",
	    "XuDaoli"],
	    "Publisher": "ituring.com.cn",
	    "IsPublished": true,
	    "Price": 9.99,
	    "Sales": 1000000
	}`)
	var r interface{}
	json.Unmarshal(b, &r)
	fmt.Println(r) // map[Authors:[XuShiwei HughLv Pandaman GuaguaSong HanTuo BertYuan XuDaoli] IsPublished:true Price:9.99 Publisher:ituring.com.cn Sales:1e+06 Title:Go语言编程]

	// 首先我们来看一下基本数据类型到 JSON 字符串的编码过程。
	// 这里是一些原子值的例子。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))       //true
	fmt.Printf("%T,%v", bolB, bolB) //[]uint8,[116 114 117 101]
	intB, _ := json.Marshal(1)
	fmt.Printf("%T,%v", intB, intB) //[]uint8,[49]
	fmt.Println(string(intB))       //1
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB)) //2.34
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB)) //"gopher"

	// 这里是一些切片和 map 编码成 JSON 数组和对象的例子。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))       //["apple","peach","pear"]
	fmt.Printf("%T,%v", slcB, slcB) //[]uint8,[91 34 97 112 112 108 101 34 44 34 112 101 97 99 104 34 44 34 112 101 97 114 34 93]

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))       //{"apple":5,"lettuce":7}
	fmt.Printf("%T,%v", mapB, mapB) //[]uint8,[123 34 80 97 103 101 34 58 49 44 34 70 114 117 105 116 115 34 58 91 34 97 112 112 108 101 34 44 34 112 101 97 99 104 34 44 34 112 101 97 114 34 93 125]

	// JSON 包可以自动的编码你的自定义类型。
	// 编码仅输出可导出的字段，并且默认使用他们的名字作为 JSON 数据的键。
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))        //{"Page":1,"Fruits":["apple","peach","pear"]}
	fmt.Printf("%T,%v", res1B, res1B) //[]uint8,[123 34 80 97 103 101 34 58 49 44 34 70 114 117 105 116 115 34 58 91 34 97 112 112 108 101 34 44 34 112 101 97 99 104 34 44 34 112 101 97 114 34 93 125]

	// 你可以给结构字段声明标签来自定义编码的 JSON 数据键名称。
	// 在上面 Response2 的定义可以作为这个标签这个的一个例子。
	res2D := Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) //{"page":1,"fruits":["apple","peach","pear"]}

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
	fmt.Printf("%T,%v", dat["num"], dat["num"]) //float64,6.13
	num := dat["num"].(float64)
	fmt.Printf("%T,%v", num, num) //float64,6.13
	// fmt.Println(num)

	// 访问嵌套的值需要一系列的转化。
	fmt.Printf("%T,%v", dat["strs"], dat["strs"]) //[]interface {},[a b]
	strs := dat["strs"].([]interface{})
	fmt.Printf("%T,%v", strs, strs)       //[]interface {},[a b]
	fmt.Printf("%T,%v", strs[0], strs[0]) //string,a
	// fmt.Println(dat["strs"][0]) //dat["strs"][0] (type interface {} does not support indexing)
	str1 := strs[0].(string)
	fmt.Println(str1) //a
	os.Exit(2)

	// 我们也可以解码 JSON 值到自定义类型。
	// 这个功能的好处就是可以为我们的程序带来额外的类型安全加强，并且消除在访问数据时的类型断言。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)           //&{1 [apple peach]}
	fmt.Println(res.Fruits[0]) //apple

	// 在上面的例子中，我们经常使用 byte 和 string 作为使用标准输出时数据和 JSON 表示之间的中间值。
	// 我们也可以和os.Stdout 一样，直接将 JSON 编码直接输出至 os.Writer流中，或者作为 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k, v1 := range v {
			if k != "Title" {
				v[k] = v1
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
