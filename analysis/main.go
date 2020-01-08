package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

//保存命令行参数
type cmdParams struct {
	logFilePath string
	routineNum  int
}

//pv，uv通道使用的数据结构
type urlData struct {
	uid  string  //urlData唯一标识
	data digData //内容
}

type digData struct {
	time  string
	url   string
	refer string
	ua    string
}

type urlNode struct {
}

type storageBlock struct {
	counterType  string  `description:"pv、uv等类型"`
	storageModel string  `description:"存储类型，存储格式"`
	unode        urlNode `description:"存储内容"`
}

const REQUEST = "req:{"
const RESPONSE = "req:&"

var log = logrus.New()

func init() {
	log.Out = os.Stdout //日志打印到标准输出
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	//	获取参数
	logFilePath := flag.String("logFilePath", "D:/workspace/go_workspace/go_practise/simple.log", "源文件")
	routineNum := flag.Int("routineNum", 5, "")
	l := flag.String("l", "/tmp/log", "日志文件")
	flag.Parse()
	params := cmdParams{*logFilePath, *routineNum}

	//打日志
	logFd, err := os.OpenFile(*l, os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		log.Out = logFd //日志打印到日志文件
		defer logFd.Close()
	}
	log.Infoln("Exec start")
	log.Infof("Params:logPath:%v,routineNum:%v", params.logFilePath, params.routineNum)

	//初始化一些channel,用于数据传递
	var logChannel = make(chan string, 3*params.routineNum)
	var pvChannel = make(chan urlData, params.routineNum)
	var uvChannel = make(chan urlData, params.routineNum)
	var storageChannel = make(chan storageBlock, params.routineNum)

	//日志消费者
	go readFileLineByLine(params, logChannel)

	//创建一组日志处理
	for i := 0; i < params.routineNum; i++ {
		go logConsumer(logChannel, pvChannel, uvChannel) //从logChannel读取数据，解析后放进pvChannel,uvChannel
	}

	//创建PV,UV统计器，如果有其他统计，可以如下添加
	go uvCounter(pvChannel, storageChannel)
	go pvCounter(uvChannel, storageChannel)

	//创建存储器
	go dataStorage(storageChannel)

	time.Sleep(time.Second * 5)
}

func readFileLineByLine(params cmdParams, logChannel chan string) error {
	file, err := os.Open(params.logFilePath)
	if err != nil {
		log.Warningf("open file error:%s\n", err)
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	count := 0 //没1000行标识一次
	for {
		line, err := reader.ReadString('\n') //byte
		count++
		if count%(1000*params.routineNum) == 0 {
			log.Infof("has read line:%v\n", count)
		}

		if err != nil {
			if err == io.EOF {
				log.Infof("finished reading log file:%s\n", err)
			} else {
				log.Infof("reading log file error:%s\n", err)
			}
			return err
		}
		logChannel <- line
	}
	return nil
}

func logConsumer(logChannel chan string, pvChannel chan urlData, uvChannel chan urlData) {
	for line := range logChannel {
		data := cutLogFetchData(line)

		//	uid
		hasher := md5.New()
		hasher.Write([]byte(data.ua + data.refer))
		uid := hex.EncodeToString(hasher.Sum(nil))

		udata := urlData{uid, data}
	}
}

func cutLogFetchData(line string) digData {
	dig := digData{}
	line = strings.Trim(line, "")
	if i := strings.Index(line, REQUEST); i != -1 {

	} else if i := strings.Index(line, RESPONSE); i != -1 {

	}

	return dig
}

func uvCounter(pvChannel chan urlData, storageChannel chan storageBlock) {

}
func pvCounter(uvChannel chan urlData, storageChannel chan storageBlock) {

}

func dataStorage(storageChannel chan storageBlock) {

}
