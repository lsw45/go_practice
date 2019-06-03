/*
需求描述

申请建群：向服务端发送聊天消息“建群#群昵称”

服务端审核：如果没有同名群存在，就返回“建群成功”否则返回“群已存在”

创建群结构体：属性包括群主、群昵称、群成员；

查看群信息：任何人可以通过向服务端发送聊天信息“群信息#所有”或者“群信息#昵称”查看群信息

服务端返回：服务端返回格式

 群昵称：xxx
 群主：xxx
 群人数：xxx
申请加群：任何人可以通过向服务端发送聊天信息“加群#群昵称”申请加群

加群审核：服务端将加群信息转发给群主，由群主确认，有服务端转发结果

群聊：群成员向服务端发送消息“群信息-群昵称#信息内容”，服务端将群信息以“成员昵称：信息内容”广播给群里所有人

实现文件上传
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"
)

var (
	//客户端信息,用昵称为键
	//allClientsMap = make(map[string]net.Conn)
	allClientsMap = make(map[string]*Client)

	//所有群
	allGroupsMap map[string]*Group

	//basePath
	basePath = "C:/Users/sirouyang/Desktop/W4/day5/02飞狗聊天/server/uploads/"
)

func init() {
	allGroupsMap = make(map[string]*Group)
	allGroupsMap["示例群"] = NewGroup("示例群", &Client{name: "系统管理员"})
}

func SHandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
		os.Exit(1)
	}
}

func main() {

	//建立服务端监听
	listener, e := net.Listen("tcp", "127.0.0.1:8888")
	SHandleError(e, "net.Listen")
	defer func() {
		for _, client := range allClientsMap {
			client.conn.Write([]byte("all:服务器进入维护状态，大家都洗洗睡吧！"))
		}
		listener.Close()
	}()

	for {
		//循环接入所有女朋友
		conn, e := listener.Accept()
		SHandleError(e, "listener.Accept")
		clientAddr := conn.RemoteAddr()

		//TODO:接收并保存昵称
		buffer := make([]byte, 1024)
		var clientName string
		for {
			n, err := conn.Read(buffer)
			SHandleError(err, "conn.Read(buffer)")
			if n > 0 {
				clientName = string(buffer[:n])
				break
			}
		}
		fmt.Println(clientName + "上线了")

		//TODO:将每一个女朋友丢入map
		client := &Client{conn, clientName, clientAddr.String()}
		allClientsMap[clientName] = client

		//TODO:给已经在线的用户发送上线通知——使用昵称
		for _, client := range allClientsMap {
			client.conn.Write([]byte(clientName + "上线了"))
		}

		//在单独的协程中与每一个具体的女朋友聊天
		go ioWithClient(client)
	}

}

//与一个Client做IO
func ioWithClient(client *Client) {
	//clientAddr := conn.RemoteAddr().String()
	buffer := make([]byte, 1024)

	for {
		n, err := client.conn.Read(buffer)
		if err != io.EOF {
			SHandleError(err, "conn.Read")
		}

		if n > 0 {
			msgBytes := buffer[:n]
			if bytes.Index(msgBytes, []byte("upload")) == 0 {
				/*处理文件上传*/

				//拿到数据包头（文件名）
				msgStr := string(msgBytes[:100])
				fileName := strings.Split(msgStr, "#")[1]

				//拿到数据包身体（文件字节）
				fileBytes := msgBytes[100:]

				//将文件字节写入指定位置
				err := ioutil.WriteFile(basePath+fileName, fileBytes, 0666)
				SHandleError(err, "ioutil.WriteFile")
				fmt.Println("文件上传成功")
				SendMsg2Client("文件上传成功", client)

			} else {
				/*处理字符消息*/
				//拿到客户端消息
				msg := string(msgBytes)
				fmt.Printf("%s:%s\n", client.name, msg)

				//将客户端说的每一句话记录在【以他的名字命名的文件里】
				writeMsgToLog(msg, client)

				strs := strings.Split(msg, "#")
				if len(strs) > 1 {
					//要发送的目标昵称
					header := strs[0]
					body := strs[1]

					switch header {

					//世界消息
					case "all":
						handleWorldMsg(client, body)

						//建群申请
					case "group_setup":
						handleGroupSetup(body, client)

						//查看群信息
					case "group_info":
						handleGroupInfo(body, client)

						//加群申请
					case "group_join":
						group, ok := allGroupsMap[body]
						//如果群不存在
						if !ok {
							SendMsg2Client("查无此群,fuckoff", client)
							continue
						}

						//发出加群申请
						SendMsg2Client(client.name+"申请加入群"+body+",是否同意？", group.Owner)
						SendMsg2Client("申请已发送，请等待群主审核", client)

						//处理群主的回复
					case "group_joinreply":
						//group_joinreply#no@zhangsan@东方艺术殿堂交流群

						//拿到回复、申请人昵称、群昵称、
						strs := strings.Split(body, "@")
						answer := strs[0]
						applicantName := strs[1]
						groupName := strs[2]

						//判断是否群昵称和申请人是否合法
						group, ok1 := allGroupsMap[groupName]
						toWhom, ok2 := allClientsMap[applicantName]

						//自动执行加群申请
						if ok1 && ok2 {
							NewGroupJoinReply(client, toWhom, group, answer).AutoRun()
						}

					default:
						//点对点消息
						handleP2PMsg(header, client, body)
					}

				} else {

					//客户端主动下线
					if msg == "exit" {
						//将当前客户端从在线用户中除名
						//向其他用户发送下线通知
						for name, c := range allClientsMap {
							if c == client {
								delete(allClientsMap, name)
							} else {
								c.conn.Write([]byte(name + "下线了"))
							}
						}
					} else if strings.Index(msg, "log@") == 0 {
						//log@all
						//log@张全蛋
						filterName := strings.Split(msg, "@")[1]
						//向客户端发送它的聊天日志
						go sendLog2Client(client, filterName)
					} else {
						client.conn.Write([]byte("已阅：" + msg))
					}

				}
			}

		}
	}

}

/*处理点对点消息*/
func handleP2PMsg(header string, client *Client, body string) {
	for key, c := range allClientsMap {
		if key == header {
			c.conn.Write([]byte(client.name + ":" + body))

			//在点对点消息的目标端也记录日志
			go writeMsgToLog(client.name+":"+body, c)
			break
		}
	}
}

/*处理查看群信息*/
func handleGroupInfo(body string, client *Client) {
	if body == "all" {
		//查看所有群信息
		info := ""
		for _, group := range allGroupsMap {
			info += group.String() + "\n"
		}
		SendMsg2Client(info, client)
	} else {
		//查看单个群信息
		if group, ok := allGroupsMap[body]; ok {
			SendMsg2Client(group.String(), client)
		} else {
			SendMsg2Client("查无此群,stupid!", client)
		}

	}
}

/*处理建群申请*/
func handleGroupSetup(body string, client *Client) {
	if _, ok := allGroupsMap[body]; !ok {
		//建群
		newGroup := NewGroup(body, client)

		//将新群添加到所有群集合
		allGroupsMap[body] = newGroup

		//通知群主建群成功
		SendMsg2Client("建群成功", client)
	} else {
		//要创建的群已经存在
		SendMsg2Client("要创建的群已经存在", client)
	}
}

/*处理世界消息*/
func handleWorldMsg(client *Client, body string) {
	for _, c := range allClientsMap {
		c.conn.Write([]byte(client.name + ":" + body))
	}
}

func SendMsg2Client(msg string, client *Client) {
	client.conn.Write([]byte(msg))
}

//向客户端发送它的聊天日志
func sendLog2Client(client *Client, filterName string) {
	//读取聊天日志
	logBytes, e := ioutil.ReadFile("D:/BJBlockChain1801/demos/W4/day1/01ChatRoomII/logs/" + client.name + ".log")
	SHandleError(e, "ioutil.ReadFile")

	if filterName != "all" {
		//查找与某个人的聊天记录
		//从内容中筛选出带有【filterName#或filterName:】的行，拼接起来
		logStr := string(logBytes)
		targetStr := ""
		lineSlice := strings.Split(logStr, "\n")
		for _, lineStr := range lineSlice {
			if len(lineStr) > 20 {
				contentStr := lineStr[20:]
				if strings.Index(contentStr, filterName+"#") == 0 || strings.Index(contentStr, filterName+":") == 0 {
					targetStr += lineStr + "\n"
				}
			}
		}
		client.conn.Write([]byte(targetStr))
	} else {
		//查询所有的聊天记录
		//向客户端发送
		client.conn.Write(logBytes)
	}

}

//将客户端说的一句话记录在【以他的名字命名的文件里】
func writeMsgToLog(msg string, client *Client) {
	//打开文件
	file, e := os.OpenFile(
		"D:/BJBlockChain1801/demos/W4/day1/01ChatRoomII/logs/"+client.name+".log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644)
	SHandleError(e, "os.OpenFile")
	defer file.Close()

	//追加这句话
	logMsg := fmt.Sprintln(time.Now().Format("2006-01-02 15:04:05"), msg)
	file.Write([]byte(logMsg))
}
