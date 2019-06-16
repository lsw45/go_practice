package main

import (
	"github.com/astaxie/beego/logs"
)

/*
* 1.同一个日志对象，只能使用一个初始化

 */
func main() {
	Console()
	//ToFile()
	//consoleAndFile()
	// conn()
	//smtp()
}

//控制台输出
func Console() {
	log := logs.NewLogger()

	log.SetLogger(logs.AdapterConsole)
	//设置打印函数及行号
	log.EnableFuncCallDepth(true)
	log.Debug("log1---> my book is bought in the year of ")
	log.Critical("log1---> oh,crash")
}

//文件输出
func ToFile() {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterFile, `{"filename":"log2.log","maxlines":1000,"maxsize":1000,"daily":true,"maxdays":10,"color":true}`)
	//l.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":1000,"maxsize":1000,"daily":true,"maxdays":10,"color":true}`)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(2)

	log.Debug("log2---> my book is bought in the year of ")
	log.Critical("log2---> oh,crash")
}

//控制台，文件同时输出
func consoleAndFile() {
	log := logs.NewLogger()

	log.SetLogger(logs.AdapterConsole)
	log.SetLogger(logs.AdapterFile, `{"filename":"log3-6.log","level":6,"maxlines":1000,"maxsize":1000,"daily":true,"maxdays":10,"color":true}`)

	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(2)
	//log.Async(1e3)

	//1-7级别递减，默认是trace，显示当前数字以前的级别，例如：3时，显示【Emergency】【Alert】【Critical】【Error】
	log.Emergency("log3--->Emergency")
	log.Alert("log3--->Alert")       //1
	log.Critical("log3--->Critical") //2
	log.Error("log3--->Error")       //3
	log.Warn("log3--->Warning")      //4
	log.Notice("log3--->Notice")     //5
	log.Info("log3--->Info")         //6
	log.Debug("log3--->Debug")       //7
	log.Trace("log3--->Trace")
}

//网络数据
func conn() {
	log := logs.NewLogger()
	//udp 传输到本地端口7020
	log.SetLogger(logs.AdapterConn, `{"net":"udp","addr":"172.20.36.141:7020"}`)

	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(2)

	log.Emergency("log4--->Emergency")
}

//邮件发送
func smtp() {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterMail, `{"username":"luciferofwg@gmail.com","password":"xxxxxxxxx","host":"imap.gmail.com:993","sendTos":["958730879@qq.com"]}`)

	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(2)
	log.Emergency("log5--->Emergency")
}
