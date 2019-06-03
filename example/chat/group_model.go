package main

import (
	"strconv"
)

/*
·创建群结构体：属性包括群主、群昵称、群成员；
*/
type Group struct {
	//群昵称
	Name string
	//群主
	Owner *Client
	//群成员
	Members []*Client
}

/*
群昵称：xxx
群主：xxx
群人数：xxx
*/
func (g *Group) String() string {
	info := "群昵称：" + g.Name + "\n"
	info += "群  主：" + g.Owner.name + "\n"
	info += "群人数：" + strconv.Itoa(len(g.Members)) + "人\n"
	return info
}

/*添加新成员*/
func (g *Group) AddClient(client *Client) {
	g.Members = append(g.Members, client)
}

/*建群工厂方法*/
func NewGroup(name string, owner *Client) *Group {
	group := new(Group)
	group.Name = name
	group.Owner = owner
	group.Members = make([]*Client, 0)
	group.Members = append(group.Members, owner)
	return group
}

/*加群申请回复*/
type GroupJoinReply struct {
	//发送人
	fromWhom *Client
	//申请人
	toWhom *Client

	//申请的群
	group *Group
	//同意与否
	answer string
}

/*加群申请工厂方法*/
func NewGroupJoinReply(fromWhom, toWhom *Client, group *Group, answer string) *GroupJoinReply {
	reply := new(GroupJoinReply)
	reply.fromWhom = fromWhom
	reply.toWhom = toWhom
	reply.group = group
	reply.answer = answer
	return reply
}

//加群审核的自动执行
func (reply *GroupJoinReply) AutoRun() {
	if reply.group.Owner == reply.fromWhom {
		//回复是群主发的
		if reply.answer == "yes" {
			reply.group.AddClient(reply.toWhom)
			SendMsg2Client("你已成功加入"+reply.group.Name, reply.toWhom)
		} else {
			SendMsg2Client(reply.group.Name+"群主已经拒绝了您的加群请求，fuckoff！", reply.toWhom)
		}
	} else {
		//不是群主发的可以将“伪群主”封号
		SendMsg2Client("根据《中华人民共和国促进装逼法》,你已获得《葵花宝典》的练习权，执法人员将送书上门并监督练习", reply.fromWhom)
	}
}
