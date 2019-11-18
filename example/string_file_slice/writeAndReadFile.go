package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

/*给原始字节切片起一个类别名*/
type Data []byte

type DataFile struct {
	//文件
	f *os.File

	//数据块长度
	datalen uint32

	//最后一次写入的偏移量（字节）
	woffset int64
	//最后一次读取的偏移量（字节）
	roffset int64

	//文件的读写锁
	fmutex sync.RWMutex

	//条件变量
	rcond *sync.Cond
}

/*
工厂方法
参数：
path string		文件路径
datalen uint32	指定数据块大小
----------
返回值：
DataFile	数据文件对象
error		错误
*/
func NewDataFile(path string, datalen uint32) (*DataFile, error) {
	//得到一个可读可写（覆盖式写入）的文件
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	/*数据块大小不能为0*/
	if datalen == 0 {
		return nil, errors.New("Invalid data legth!")
	}

	//创建指定的myDataFile对象，并指定IO文件和数据块大小
	df := &DataFile{f: f, datalen: datalen}

	//初始化df对象的条件变量
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}

/*返回下一次要读取的数据块的序列号*/
func (df *DataFile) Rsn() int64 {
	//同步加载最后一次读取的字节偏移量
	offset := atomic.LoadInt64(&df.roffset)

	//返回下一次读取的数据块的序列号
	return offset / int64(df.datalen)
}

/*返回下一次要写入的数据块的序列号*/
func (df *DataFile) Wsn() int64 {
	//同步获取当前df的写入的字节偏移量
	offset := atomic.LoadInt64(&df.woffset)

	// 返回下一次写入的数据块的序列号
	return offset / int64(df.datalen)
}

/*返回块文件的大小*/
func (df *DataFile) DataLen() uint32 {
	//return df.datalen

	//同步获取当前df的数据块大小并返回
	return atomic.LoadUint32(&df.datalen)
}

/*
可以读，返回值：
rsn 	read-serial-number 当前读取到的【数据块序列号】
data	原始字节数据
err 	错误
*/
func (df *DataFile) Read() (rsn int64, data Data, err error) {

	var offset int64
	//offset = df.roffset
	//加载最后一次读取的字节偏移量
	offset = atomic.LoadInt64(&df.roffset)

	//计算本次读取的数据块的序列号
	rsn = offset / int64(df.datalen) //第几块数据块

	//创建一个数据块大小的缓冲区
	buffer := make([]byte, df.datalen)

	//加读锁
	df.fmutex.RLock()
	fmt.Println("Read get lock")
	defer func() {
		//返回之前释放读锁
		df.fmutex.RUnlock()
		fmt.Println("Read release lock")
	}()

	for {
		//从指定的字节偏移量处进行读取
		_, err = df.f.ReadAt(buffer, offset)
		//fmt.Println("n,err=",n,e)
		if err != nil {

			//如果已经读到了文件末尾
			if err == io.EOF {
				fmt.Println("eof:Read release lock")
				//阻塞等待有新的内容写入
				df.rcond.Wait()
				fmt.Println("eof:Read get lock")
				continue
			}
			return
		}

		//正常地读到数据，并返回
		data = buffer

		//通过原子操作，让最后一次读取的字节偏移量+=df.datalen
		atomic.AddInt64(&df.roffset, int64(df.datalen))

		return
	}
}

func (df *DataFile) Write(data Data) (wsn int64, err error) {
	//获取并更新写偏移量
	var offset int64
	//offset = df.woffset
	//加载上次写入的字节偏移量
	offset = atomic.LoadInt64(&df.woffset)

	/*如果要写入的数据超过一个数据块的长度，就进行截取操作，否则直接使用*/
	var buffer []byte
	if len(data) > int(df.datalen) {
		buffer = data[0:df.datalen]
	} else {
		buffer = data
	}

	/*加写锁准备进行写入*/
	df.fmutex.Lock()
	fmt.Println("Write get lock")
	defer func() {
		//操作完毕释放写锁
		df.fmutex.Unlock()
		fmt.Println("Write release lock")
	}()

	/*进行数据写入*/
	//写入数据，并向读取协程发送通知信号
	_, err = df.f.WriteAt(buffer, offset)

	if err == nil {
		//本次写入导致最后一次写入的字节偏移量+=df.datalen
		atomic.AddInt64(&df.woffset, int64(df.datalen))

		//计算本次写入的块序列号
		wsn = offset / int64(df.datalen)

		//向读取协程发送数据更新通知
		fmt.Println("Write signal")
		df.rcond.Signal()
	}

	return
}
func main() {
	var wg sync.WaitGroup

	df, e := NewDataFile(`testFile.txt`, 15)
	fmt.Println(df, e)

	fmt.Println("数据块大小=", df.DataLen())  //3
	fmt.Println("下一次读取的块序列号=", df.Rsn()) //0
	fmt.Println("下一次写入的块序列号=", df.Wsn()) //0

	wg.Add(2)
	/*写入3个数据块*/
	go func() {
		wsn, err := df.Write(Data("明月几时有"))
		fmt.Println("wsn, err=", wsn, err)

		wsn, err = df.Write(Data("把酒问青天"))
		fmt.Println("wsn, err=", wsn, err)

		wsn, err = df.Write(Data("不知天上宫阙"))
		fmt.Println("wsn, err=", wsn, err)

		wg.Done()
	}()

	/*读取3个数据块*/
	go func() {
		rsn, data, err := df.Read()
		fmt.Println("rsn, data, err=", rsn, string(data), err)

		rsn, data, err = df.Read()
		fmt.Println("rsn, data, err=", rsn, string(data), err)

		rsn, data, err = df.Read()
		fmt.Println("rsn, data, err=", rsn, string(data), err)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main over")

}
