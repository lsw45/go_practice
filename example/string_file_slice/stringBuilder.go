package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// 示例1。
	var builder1 strings.Builder
	builder1.Grow(100)
	fmt.Printf("The first output(%d):\n%q\n", builder1.Len(), builder1.String()) //The first output(0):""
	builder1.WriteString("A Builder is used to efficiently build a string using Write methods.")
	fmt.Printf("The first output(%d):\n%q\n", builder1.Len(), builder1.String()) //The first output(68):"A Builder is used to efficiently build a string using Write methods."
	fmt.Println()
	builder1.WriteByte(' ')
	builder1.WriteString("It minimizes memory copying. The zero value is ready to use.")
	builder1.Write([]byte{'\n', '\n'})
	builder1.WriteString("Do not copy a non-zero Builder.")
	fmt.Printf("The second output(%d):\n\"%s\"\n", builder1.Len(), builder1.String()) //The second output(162):
	//"A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use.
	//Do not copy a non-zero Builder."
	fmt.Println()

	// 示例2。
	fmt.Println("Grow the builder ...")
	builder1.Grow(10)
	builder1.Grow(200)
	fmt.Printf("The length of contents in the builder is %d.\n", builder1.Len()) //The length of contents in the builder is 162.
	fmt.Println()

	// 示例3。
	fmt.Println("Reset the builder ...")
	builder1.Reset()
	fmt.Printf("The third output(%d):\n%q\n", builder1.Len(), builder1.String()) //The third output(0):""

	stringReader()
}

func stringReader() {
	// 示例1。
	reader1 := strings.NewReader(
		"NewReader returns a new Reader reading from s. " +
			"It is similar to bytes.NewBufferString but more efficient and read-only.")
	fmt.Printf("The size of reader: %d\n", reader1.Size()) //The size of reader: 119
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len())) //The reading index in reader: 0

	buf1 := make([]byte, 47)
	n, _ := reader1.Read(buf1)
	fmt.Printf("%d bytes were read. (call Read)\n", n) //47 bytes were read. (call Read)
	fmt.Printf("%q", string(buf1))                     //NewReader returns a new Reader reading from s.
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len())) //The reading index in reader: 47
	fmt.Println()

	// 示例2。
	buf2 := make([]byte, 21)
	offset1 := int64(64)
	n, _ = reader1.ReadAt(buf2, offset1)
	fmt.Printf("%q", string(buf2))                                            //bytes.NewBufferString
	fmt.Printf("%d bytes were read. (call ReadAt, offset: %d)\n", n, offset1) //21 bytes were read. (call ReadAt, offset: 64)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len())) //The reading index in reader: 47
	fmt.Println()

	// 示例3。
	offset2 := int64(17)
	expectedIndex := reader1.Size() - int64(reader1.Len()) + offset2
	fmt.Printf("Seek with offset %d and whence %d ...\n", offset2, io.SeekCurrent) //Seek with offset 17 and whence 1 ...
	readingIndex, _ := reader1.Seek(offset2, io.SeekCurrent)
	fmt.Printf("The reading index in reader: %d (returned by Seek)\n", readingIndex) //The reading index in reader: 64 (returned by Seek)
	fmt.Printf("The reading index in reader: %d (computed by me)\n", expectedIndex)  //The reading index in reader: 64 (computed by me)

	n, _ = reader1.Read(buf2)
	fmt.Printf("%d bytes were read. (call Read)\n", n) //21 bytes were read. (call Read)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len())) //The reading index in reader: 85
}
