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
