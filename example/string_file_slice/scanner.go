package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 对 os.Stdin 使用一个带缓冲的 scanner，让我们可以直接使用方便的 Scan 方法来直接读取一行，每次调用该方法可以让 scanner 读取下一行。
	// bufio.NewScanner的参数是*os.File类型
	scanner := bufio.NewScanner(os.Stdin)

	// Text 返回当前的 token，现在是输入的下一行。
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		// 写出大写的行。
		fmt.Println(ucl)
	}

	// 检查 Scan 的错误。
	// 文件结束符是可以接受的，并且不会被 Scan 当作一个错误。
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
