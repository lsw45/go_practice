package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// TestPeek()
	// TestRead()
	TestBuffered()
	// TestReadByte()
	// TestUnreadByte()
	// TestReadRune()
	// TestUnReadRune()
	// TestReadLine()
	// TestReadBytes()
	// TestReadString()
	// TestWriteTo()
	// TestNewWriter()
	// TestWrite()
	// TestWriteString()
	// TestWriteByte()
	// TestWriteRune()
	// TestReadFrom()
	// TestReadWriter()
	// TestNewScanner()
	// TestSplit()
	// TestScan()
	// TestScanBytes()
	// TestScanRunes()
	// TestScanWords()
	// TestScanLines()
}

func TestPeek() {
	/*
	   func NewReaderSize(rd io.Reader, size int) *Reader
	   NewReaderSize 将 rd 封装成一个带缓存的 bufio.Reader 对象，
	   缓存大小由 size 指定（如果小于 16 则会被设置为 16）。
	   minReadBufferSize = 16
	   如果 rd 的基类型就是有足够缓存的 bufio.Reader 类型，则直接将
	   rd 转换为基类型返回。

	   NewReader()方法返回一个默认大小的带缓存的bufio.Reader对象
	   即 NewReaderSize(rd, 4096)
	*/
	s := strings.NewReader("hello world")

	// func (b *Reader) Reset(r io.Reader)
	// Reset丢弃缓冲中的数据，清除任何错误，将b重设为其下层从r读取数据。
	s.Reset("my name is card")

	// func (b *Reader) Peek(n int) ([]byte, error)
	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 个字节的数据，
	// 该操作不会将数据读出，只是引用，引用的数据在下一次读取操作之
	// 前是有效的。如果切片长度小于 n，则返回一个错误信息说明原因。
	// 如果 n 大于缓存的总大小，则返回 ErrBufferFull。
	br := bufio.NewReader(s)

	b, _ := br.Peek(5)
	b[0] = 'a'
	b, _ = br.Peek(5)
	fmt.Printf("%q\n", b) // "ay na"
}

func TestRead() {
	// Read 从 b 中读出数据到 p 中，返回写入p的字节数
	// 读取到达结尾时，返回值n将为0而err将为io.EOF。
	// 如果缓存不为空，则只能读出缓存中的数据，不会从底层 io.Reader
	// 中提取数据，如果缓存为空，则：
	// 1、len(p) >= 缓存大小，则跳过缓存，直接从底层 io.Reader 中读
	// 出到 p 中。
	// 2、len(p) < 缓存大小，则先将数据从底层 io.Reader 中读取到缓存
	// 中，再从缓存读取到 p 中。
	// func (b *Reader) Read(p []byte) (n int, err error)
	s := strings.NewReader("123456789")
	br := bufio.NewReader(s)
	b := make([]byte, 4)
	n, err := br.Read(b)
	fmt.Printf("%s %v %v\n", b[:n], n, err) // 1234 4

	n, err = br.Read(b)
	fmt.Printf("%s %v %v\n", b[:n], n, err) // 5678 4

	n, err = br.Read(b)
	fmt.Printf("%s %v %v\n", b[:n], n, err) // 9 1

	n, err = br.Read(b)
	fmt.Printf("%s %v %v\n", b[:n], n, err) // 0 EOF
}

func TestBuffered() {
	// 返回可以从缓存中读取的字节数
	// func (b *Reader) Buffered() int { return b.w - b.r }
	s := strings.NewReader("123456789")
	br := bufio.NewReader(s)
	b := make([]byte, 3)
	br.Read(b)
	fmt.Println(br.Buffered()) // 6

	br.Read(b)
	fmt.Println(br.Buffered()) // 3
}

func TestReadByte() {
	// ReadByte读取并返回一个字节。如果没有可用的数据，会返回错误。
	// func (b *Reader) ReadByte() (c byte, err error)
	origin := "abcd"
	s := strings.NewReader(origin)
	br := bufio.NewReader(s)
	// 第一次读取
	tmp, err := br.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", tmp)    // 'a'
	fmt.Println(br.Buffered()) // 3
	for i := 0; i < len(origin); i++ {
		tmp, err = br.ReadByte()
		if err != nil {
			// panic: EOF  因为已经读取了1个字符 缓存中只剩下3个
			// 所以在读取第4个字符报错EOF
			panic(err)
		}
	}
}

func TestUnreadByte() {
	// 撤消最后读出的字节
	s := strings.NewReader("abcde")
	br := bufio.NewReader(s)
	tmp, _ := br.ReadByte()
	fmt.Printf("%q\n", tmp)    // 'a'
	fmt.Println(br.Buffered()) // 4
	br.UnreadByte()            // 撤销吐出,即栈中弹出的a元素又放回来了
	fmt.Println(br.Buffered()) // 5
	tmp, _ = br.ReadByte()
	fmt.Printf("%q\n", tmp) // 'a'
}

func TestReadRune() {
	// ReadRune读取一个utf-8编码的unicode码值
	chinese := "中国人"
	s := strings.NewReader(chinese)
	br := bufio.NewReader(s)
	tmp, _, err := br.ReadRune()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", tmp) // '中'
}

func TestUnReadRune() {
	chinese := "中国人"
	s := strings.NewReader(chinese)
	br := bufio.NewReader(s)
	tmp, _, err := br.ReadRune()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", tmp) // '中'
	br.UnreadRune()
	tmp, _, err = br.ReadRune()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", tmp) // '中'
}

func TestReadLine() {
	// ReadLine 是一个低水平的行读取原语，大多数情况下，应该使用
	// ReadBytes('\n') 或 ReadString('\n')，或者使用一个 Scanner。
	//
	// ReadLine 通过调用 ReadSlice 方法实现，返回的也是缓存的切片。用于
	// 读取一行数据，不包括行尾标记（\n 或 \r\n）。
	//
	// 只要能读出数据，err 就为 nil。如果没有数据可读，则 isPrefix 返回
	// false，err 返回 io.EOF。
	//
	// 如果找到行尾标记，则返回查找结果，isPrefix 返回 false。
	// 如果未找到行尾标记，则：
	// 1、缓存不满，则将缓存填满后再次查找。
	// 2、缓存是满的，则返回整个缓存，isPrefix 返回 true。
	//
	// 整个数据尾部“有一个换行标记”和“没有换行标记”的读取结果是一样。
	//
	// 如果 ReadLine 读取到换行标记，则调用 UnreadByte 撤销的是换行标记，
	// 而不是返回的数据。
	// func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)

	s := strings.NewReader("123\nzzz")
	br := bufio.NewReader(s)
	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = br.ReadLine()
		// "123" false
		// "zzz" false
		// "" false EOF
		fmt.Printf("%q %t %v\n", line, isPrefix, err)
	}
}

func TestReadSlice() {
	// ReadSlice 在 b 中查找 delim 并返回 delim 及其之前的所有数据。
	// 该操作会读出数据，返回的切片是已读出的数据的引用，切片中的数据
	// 在下一次读取操作之前是有效的。
	//
	// 如果找到 delim，则返回查找结果，err 返回 nil。
	// 如果未找到 delim，则：
	// 1、缓存不满，则将缓存填满后再次查找。
	// 2、缓存是满的，则返回整个缓存，err 返回 ErrBufferFull。
	//
	// 如果未找到 delim 且遇到错误（通常是 io.EOF），则返回缓存中的所
	// 有数据和遇到的错误。
	//
	// 因为返回的数据有可能被下一次的读写操作修改，所以大多数操作应该
	// 使用 ReadBytes 或 ReadString，它们返回的是数据的拷贝。
	// func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
	s := strings.NewReader("ABC DEF GHI")
	br := bufio.NewReader(s)

	w, err := br.ReadSlice(' ')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", w) // "ABC "

	w, err = br.ReadSlice(' ')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", w) // "DEF "

	w, err = br.ReadSlice(' ')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", w) // panic: EOF
}

func TestReadBytes() {
	// ReadBytes 功能同 ReadSlice，只不过返回的是缓存的拷贝。
	// func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
	s := strings.NewReader("ABC,EFG,HIJ")
	br := bufio.NewReader(s)
	line, err := br.ReadBytes(',')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", line) // "ABC,"

	line, err = br.ReadBytes(',')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", line) // "EFG,"

	line, err = br.ReadBytes(',')
	if err != nil {
		panic(err) // panic: EOF
	}
	fmt.Printf("%q\n", line)

}

func TestReadString() {
	// ReadString 功能同 ReadBytes，只不过返回的是字符串。
	// func (b *Reader) ReadString(delim byte) (line string, err error)
	s := strings.NewReader("你好,我是卡牌")
	br := bufio.NewReader(s)
	line, err := br.ReadString(',')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", line) // 你好,

	line, err = br.ReadString(',')
	if err != nil {
		panic(err) // panic: EOF
	}
	fmt.Printf("%s\n", line)
}

func TestWriteTo() {
	// WriteTo方法实现了io.WriterTo接口。
	// func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)
	b := bytes.NewBuffer(make([]byte, 0))
	br.WriteTo(b)
	fmt.Printf("%q\n", b) // "ABCDEFG"
}

// Writer实现了为io.Writer接口对象提供缓冲。
// 如果在向一个Writer类型值写入时遇到了错误，
// 该对象将不再接受任何数据，返回该错误
// 数据都写入后，调用者有义务调用Flush方法,
// 保证所有的数据都交给了下层的io.Writer。
func TestNewWriter() {

	// NewWriter创建一个具有默认大小缓冲、写入w的*Writer。
	// 相当于 NewWriterSize(wr, 4096)
	// func NewWriter(w io.Writer) *Writer

	// Buffered()返回缓冲中已使用的字节数。
	// func (b *Writer) Buffered() int

	// Available()返回缓冲中还有多少字节未使用。
	// func (b *Writer) Available() int

	// Reset丢弃缓冲中的数据，清除任何错误，将b重设为将其输出写入w。
	// func (b *Writer) Reset(w io.Writer)

	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)

	fmt.Println(bw.Available(), bw.Buffered()) // 4096  0
	bw.WriteString("card")
	fmt.Println(bw.Available(), bw.Buffered()) // 4092  4

	bw.Reset(b)
	fmt.Println(bw.Available(), bw.Buffered()) // 4096  0
}

func TestWrite() {
	// Write 将 p 中的数据写入 b 中，返回写入的字节数
	// 如果写入的字节数小于 p 的长度，则返回一个错误信息
	// func (b *Writer) Write(p []byte) (nn int, err error)

	// Flush 将缓存中的数据提交到底层的 io.Writer 中
	// func (b *Writer) Flush() error
	p := [...]byte{'a', 'b', 'c'}
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.Write(p[:])
	bw.Flush()
	fmt.Printf("%q\n", b)
}

func TestWriteString() {
	// WriteString 同 Write，只不过写入的是字符串
	// func (b *Writer) WriteString(s string) (int, error)
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("hello world")
	bw.Flush()
	fmt.Printf("%s\n", b)
}

func TestWriteByte() {
	// WriteByte写入单个字节。
	// func (b *Writer) WriteByte(c byte) error
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteByte('c')
	bw.Flush()
	fmt.Println(b)
}

func TestWriteRune() {
	// WriteRune写入一个unicode码值（的utf-8编码），返回写入的字节数和可能的错误。
	// func (b *Writer) WriteRune(r rune) (size int, err error)
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	size, err := bw.WriteRune('周')
	if err != nil {
		panic(err)
	}
	fmt.Println(size) // 3
	bw.Flush()
	fmt.Println(b) // 周
}

func TestReadFrom() {
	// ReadFrom实现了io.ReaderFrom接口。
	// func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
	// ReadFrom无需使用Flush
	b := bytes.NewBuffer(make([]byte, 0))
	s := strings.NewReader("hello world")
	bw := bufio.NewWriter(b)
	bw.ReadFrom(s)
	fmt.Println(b)
}

func TestReadWriter() {
	// ReadWriter类型保管了指向Reader和Writer类型的指针
	// 实现了io.ReadWriter接口。

	// NewReadWriter 生成bufio.ReadWriter对象
	// func NewReadWriter(r *Reader, w *Writer) *ReadWriter
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	s := strings.NewReader("hello world")
	br := bufio.NewReader(s)
	rw := bufio.NewReadWriter(br, bw)

	word, err := rw.ReadString(' ')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", word) // hello

	_, err = rw.WriteString(",I'm coming")
	if err != nil {
		panic(err)
	}
	rw.Flush()
	fmt.Println(b)
}

func TestNewScanner() {
	// Scanner 提供了一个方便的接口来读取数据，例如遍历多行文本中的行。Scan 方法会通过
	// 一个“匹配函数”读取数据中符合要求的部分，跳过不符合要求的部分。“匹配函数”由调
	// 用者指定。本包中提供的匹配函数有“行匹配函数”、“字节匹配函数”、“字符匹配函数”
	// 和“单词匹配函数”，用户也可以自定义“匹配函数”。默认的“匹配函数”为“行匹配函
	// 数”，用于获取数据中的一行内容（不包括行尾标记）
	//
	// Scanner 使用了缓存，所以匹配部分的长度不能超出缓存的容量。默认缓存容量为 4096 -
	// bufio.MaxScanTokenSize，用户可以通过 Buffer 方法指定自定义缓存及其最大容量。
	//
	// Scan 在遇到下面的情况时会终止扫描并返回 false（扫描一旦终止，将无法再继续）：
	// 1、遇到 io.EOF
	// 2、遇到读写错误
	// 3、“匹配部分”的长度超过了缓存的长度
	//
	// 如果需要对错误进行更多的控制，
	// 或“匹配部分”超出缓存容量，或需要连续扫描，
	// 则应该使用 bufio.Reader
	// func NewScanner(r io.Reader) *Scanner

	// Bytes方法返回最近一次Scan调用生成的token。
	// 底层数组指向的数据可能会被下一次Scan的调用重写。
	// func (s *Scanner) Bytes() []byte

	// Buffer()方法设置扫描时使用的初始缓冲区和最大值
	// 默认情况下，Scan使用内部缓冲区并设置MaxScanTokenSize的最大令牌大小
	s := strings.NewReader("周起\n卡牌\n程序员\n")
	bs := bufio.NewScanner(s)
	bs.Buffer(make([]byte, 0), bufio.MaxScanTokenSize)
	for bs.Scan() {
		// 周起
		// 卡牌
		// 程序员
		fmt.Printf("%s\n", bs.Bytes())
	}
}

func TestSplit() {
	// Split设置该Scanner的分割函数。默认设置为 bufio.ScanLines()
	// 本方法必须在Scan之前调用。
	// func (s *Scanner) Split(split SplitFunc)
	s := strings.NewReader("周起 卡牌 程序员")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)

	// Text返回由Scan调用生成的最新标记，
	// 作为保存其字节的新分配字符串。

	for bs.Scan() {
		fmt.Printf("%s\n", bs.Text())
	}
}

func TestScan() {
	// Scan方法获取当前位置的token（该token可以通过Bytes或Text方法获得），
	// 并让Scanner的扫描位置移动到下一个token。
	// 当扫描因为抵达输入流结尾或者遇到错误而停止时，
	// 本方法会返回false。在Scan方法返回false后，
	// Err方法将返回扫描时遇到的任何错误；
	// 除非是io.EOF，此时Err会返回nil。
	// func (s *Scanner) Scan() bool
	s := strings.NewReader("周起 卡牌 程序员")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	for bs.Scan() {
		fmt.Printf("%s %s\n", bs.Text(), bs.Bytes())
	}
}

func TestScanBytes() {
	// Bytes方法返回最近一次Scan调用生成的token。
	// 底层数组指向的数据可能会被下一次Scan的调用重写。
	s := strings.NewReader("abcd")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanBytes)
	for bs.Scan() {
		// a
		// b
		// c
		// d
		fmt.Printf("%s\n", bs.Bytes())
	}
}

func TestScanRunes() {
	// ScanRunes是用于Scanner类型的分割函数（符合SplitFunc），
	// 本函数会将每个utf-8编码的unicode码值作为一个token返回。
	s := strings.NewReader("周起卡牌程序员")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanRunes)
	for bs.Scan() {
		// 周
		// 起
		// 卡
		// 牌
		// 程
		// 序
		// 员
		fmt.Printf("%s\n", bs.Text())
	}
}

func TestScanWords() {
	// ScanRunes是用于Scanner类型的分割函数(符合SplitFunc)，
	// 本函数会将空白(参见unicode.IsSpace)
	// 分隔的片段（去掉前后空白后）作为一个token返回。
	// 本函数永远不会返回空字符串。
	s := strings.NewReader("我 是 卡 牌")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	for bs.Scan() {
		// 我
		// 是
		// 卡
		// 牌
		fmt.Printf("%s\n", bs.Text())
	}
}

func TestScanLines() {
	// 将每一行文本去掉末尾的换行标记作为一个token返回
	// 此函数的bs.Scan()的默认值
	s := strings.NewReader("卡牌\n周起\n程序员\n")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanLines)
	for bs.Scan() {
		// 卡牌
		// 周起
		// 程序员
		fmt.Printf("%s\n", bs.Text())
	}
}
