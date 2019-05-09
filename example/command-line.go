package main

import "os"
import "fmt"
import (
	"flag"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	flag.Parse()
	fmt.Println("word:", *wordPtr)

	flags()
	return

	// os.Args 提供原始命令行参数访问功能。
	// 注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 你可以使用标准的索引位置方式取得单个参数的值。
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

func flags() {
	// 返回一个相应类型的指针，
	// wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", true, "a bool")

	// 用程序中已有的参数来声明一个标志也是可以的。注意在标志声明函数中需要使用该参数的指针。
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	// 指针
	// fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
