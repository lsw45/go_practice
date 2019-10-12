package main

import (
	//"helloworld/hello"
	// "bufio"
	"encoding/json"
	// "fmt"
	"log"
	"testing"
	// "os"
	// "strings"
)

type QueryMemberResp struct {
	MemberInfo
	CardUrl     string
	EncryCardNo string // yamada加密后的卡号
}

type MemberInfo struct {
	CardNo      string
	MobileNo    string   // 手机号
	Bonus       int64    // 积分
	Level       string   // 等级
	Balance     int64    // 余额
	GiftBalance int64    // 赠送的余额
	CardTypes   []string // 卡类型
}

func TestJson(t *testing.T) {
	resp := &QueryMemberResp{
		CardUrl: "http://vka.ivp.net",
	}

	// 两种方式
	resp.MemberInfo.MobileNo = "xxxxxxxxxxxxx"
	resp.MobileNo = "199950303"
	// {"CardNo":"","MobileNo":"199950303","Bonus":0,"Level":"","Balance":0,"GiftBalance":0,"CardTypes":null,"CardUrl":"http://vka.ivp.net","EncryCardNo":""}

	str, _ := json.Marshal(resp)
	log.Println(string(str))
}
