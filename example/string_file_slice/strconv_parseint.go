package main

// 内置的 strconv 包提供了数字解析功能。
import "strconv"
import "fmt"

func main() {

	// 使用 ParseFloat 解析浮点数，这里的 64 表示表示解析的数的位数。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 在使用 ParseInt 解析整形数时，例子中的参数 0 表示自动推断字符串所表示的数字的进制。
	// 返回结果的bit大小 也就是int8 int16 int32 int64。
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	// ParseInt 会自动识别出十六进制数。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// ParseUint 也是可用的。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi 是一个基础的 10 进制整型数转换函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 在输入错误时，解析函数会返回一个错误。
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

/*
#string到int
int,err := strconv.Atoi(string)

#string到int64
int64, err := strconv.ParseInt(string, 10, 64)
//第二个参数为基数（2~36），
//第三个参数位大小表示期望转换的结果类型，其值可以为0, 8, 16, 32和64，
//分别对应 int, int8, int16, int32和int64

#int到string
string := strconv.Itoa(int)
//等价于
string := strconv.FormatInt(int64(int),10)

#int64到string
string := strconv.FormatInt(int64,10)
//第二个参数为基数，可选2~36
//对于无符号整形，可以使用FormatUint(i uint64, base int)

#float到string
string := strconv.FormatFloat(float32,'E',-1,32)
string := strconv.FormatFloat(float64,'E',-1,64)
// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'f' (-ddd.dddd，没有指数)
// 'g' ('e':大指数，'f':其它情况)
// 'G' ('E':大指数，'f':其它情况)

#string到float64
float,err := strconv.ParseFloat(string,64)

#string到float32
float,err := strconv.ParseFloat(string,32)

#int到int64
int64_ := int64(1234)
*/
