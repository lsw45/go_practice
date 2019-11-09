package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	// 需要处理的字符串
	message := "Away from keyboard. https://golang.org/"

	// 编码消息
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))

	// 输出编码完成的消息
	fmt.Println(encodedMessage)

	// 解码消息
	data, err := base64.StdEncoding.DecodeString(encodedMessage)

	// 出错处理
	if err != nil {
		fmt.Println(err)
	} else {
		// 打印解码完成的数据
		fmt.Println(string(data))
	}

	encodeAndDecode()
}

func encodeAndDecode() {
	src := []byte("Away from keyboard.")
	maxLen := base64.StdEncoding.EncodedLen(len(src))
	dst := make([]byte, maxLen)
	base64.StdEncoding.Encode(dst, src)
	fmt.Println("Encode:", dst)

	src = []byte("dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg==")
	maxLen = base64.StdEncoding.EncodedLen(len(src))
	n, err := base64.StdEncoding.Decode(dst, src)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Decode:", n)
	fmt.Println("Decode:", dst)
}
