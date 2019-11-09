package main

import (
	"flag"
	"os"
	"time"
	"github.com/sirupsen/logrus"
)

//保存命令行参数
type cmdParams struct {
	logFilePath string
	routineNum int
}

//pv，uv通道使用的数据结构
type urlData struct{
	uid string //urlData唯一标识
	data digData //内容
}

type digData struct {
	time string
	url string
	refer string
	ua string
}

type urlNode struct{

}

type storageBlock struct {
	counterType string `description:"pv、uv等类型"`
	storageModel string `description:"存储类型，存储格式"`
	unode urlNode `description:"存储内容"`
}

var log = logrus.New()
func init(){
	log.Out = os.Stdout //日志打印到标准输出
	log.SetLevel(logrus.DebugLevel)
}

func main(){
	//	获取参数
	logFilePath:=flag.String("logFilePath","/","数据存储文件")
	routineNum := flag.Int("routineNum",5,"")
	l := flag.String("l","/tmp/log","日志文件")
	flag.Parse()
	params := cmdParams{*logFilePath,*routineNum}

	//打日志
	logFd,err := os.OpenFile(*l,os.O_CREATE|os.O_WRONLY,0644)
	if err ==nil {
		log.Out = logFd //日志打印到日志文件
		defer logFd.Close()
	}
	log.Infoln("Exec start")
	log.Infof("Params:logPath:%v,routineNum:%v",params.logFilePath,params.routineNum)

	//初始化一些channel,用于数据传递
	var logChannel = make(chan string,3*params.routineNum)
	var pvChannel = make(chan urlData,params.routineNum)
	var uvChannel = make(chan urlData,params.routineNum)
	var storageChannel = make(chan storageBlock,params.routineNum)

	//日志消费者
	go readFileLineByLine(params,logChannel)

	//创建一组日志处理
	for i:=0;i<params.routineNum;i++{
		go logConsumer(logChannel,pvChannel,uvChannel) //从logChannel读取数据，解析后放进pvChannel,uvChannel
	}

	//创建PV,UV统计器，如果有其他统计，可以如下添加
	go uvCounter(pvChannel,storageChannel)
	go pvCounter(uvChannel,storageChannel)

	//创建存储器
	go dataStorage(storageChannel)

	time.Sleep(time.Second*5)
}

func readFileLineByLine(params cmdParams,logChannel chan string){

}

func logConsumer(logChannel chan string, pvChannel  chan urlData,uvChannel  chan urlData){

}

func uvCounter(pvChannel chan urlData,storageChannel chan storageBlock){

}
func pvCounter(uvChannel chan urlData,storageChannel chan storageBlock){

}

func dataStorage(storageChannel chan storageBlock){

}