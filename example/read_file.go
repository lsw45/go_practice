package main

import (
	// "bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./temp/dat.txt") //dat []byte
	check(err)
	fmt.Print(string(dat) + "\n")

	f, err := os.Open("./temp/dat.txt") // f *os.File
	check(err)

	b1 := make([]byte, 10)
	n1, err := f.Read(b1) //n1 int,b1 []byte
	check(err)
	fmt.Printf("%d bytes:%s\n", n1, string(b1))

	o2, err := f.Seek(6, 0) //从第六位开始
	check(err)
	b2 := make([]byte, 4)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d:%s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 3)
	n3, err := io.ReadAtLeast(f, b3, 3)

	check(err)
	fmt.Printf("%d bytes @ %d:%s\n", n3, o3, string(b3))

}
