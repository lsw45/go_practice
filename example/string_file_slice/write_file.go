package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	demo()
	return
	// 开始，这里是展示如写入一个字符串（或者只是一些字节）到一个文件。
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)
	// 对于更细粒度的写入，先打开一个文件。
	f, err := os.Create("/tmp/dat2")
	check(err)
	// 打开文件后，习惯立即使用 defer 调用文件的 Close操作。
	defer f.Close()
	// 你可以写入你想写入的字节切片
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
	// WriteString 也是可用的。
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)
	// 调用 Sync 来将缓冲区的信息写入磁盘。
	f.Sync()
	// bufio 提供了和我们前面看到的带缓冲的读取器一样的带缓冲的写入器。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	// 使用 Flush 来确保所有缓存的操作已写入底层写入器。
	w.Flush()
}

func demo() {
	w1 := []byte("Freedom After Success")
	err := ioutil.WriteFile("./temp/write.txt", w1, 0644)
	check(err)

	f, err := os.Create("./temp/dat.txt")
	check(err)
	defer f.Close()

	w2 := []byte{122, 3, 4, 5, 6}
	n2, err := f.Write(w2)
	check(err)
	fmt.Printf("wrote %d bytes to dat.txt\n", n2)
	n3, err := f.WriteString("writing\n")
	check(err)
	fmt.Printf("wrote %d bytes to dat.txt\n", n3)
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("bufio\n")
	fmt.Printf("wrote %d bytes to dat.txt", n4)
	w.Flush()
}
