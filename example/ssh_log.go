package main

import (
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/tealeg/xlsx"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ADDR     = "smtp.qiye.163.com:25"
	HOST     = "smtp.qiye.163.com"
	USER     = "carey.li@cardinfolink.com"
	PASSWORD = "carey.Li19940306"
)
var author = smtp.PlainAuth("", USER, PASSWORD, HOST)

func main() {
	user := "webapp"
	publicKey := "F:/讯联/SecureCrtSSH/privateKey"
	ipAdress := []string{"114.80.87.245:22","114.80.87.245:22"}
	NumRunner := len(ipAdress)
	sum := make(chan map[string]map[string]float64,NumRunner)
	summary := make(map[string]map[string]float64,0)

	var err error
	var Client *ssh.Client
	week := getWeekDay()

	for _,ip :=range ipAdress{
		go func (ip string){
			Client, err = dailPublic(user, publicKey, ip)
			if err != nil {
				fmt.Printf("连接%s失败.\n", err)
				return
			}
			defer Client.Close()

			ret := runCmd(Client,week)
			sum<-summaryAll(ret)
		}(ip)
	}

	for i:=0;i<NumRunner;i++{
		item := <-sum
		for week,val := range item{
			if len(summary[week]) == 0 {
				summary[week]= map[string]float64{"score":val["score"],"amount":val["amount"],"count":val["count"]}
			}else{
				summary[week]["score"] += val["score"]
				summary[week]["amount"] += val["amount"]
				summary[week]["count"] += val["count"]
			}
		}
	}

	//写excel
	fileName := "F:/workspace/go_example/example/file.xlsx"
	file := xlsx.NewFile()
	file,err = writeExcel(file,summary,week)
	if err != nil{
		fmt.Println(err)
	}
	err = file.Save(fileName)
	if err != nil{
		fmt.Println(err)
	}

	sendEmail(fileName)
}

func subStrValue(str,find,tag string)string{
	index := strings.Index(str,find)
	temp:=strings.Index(str[index:],tag)
	return strings.TrimRight(fmt.Sprintln(str[index+len(find):index+temp]),"\n")
}
func dail(user, password, ip_port string) (*ssh.Client, error) {
	PassWd := []ssh.AuthMethod{ssh.Password(password)}
	Conf := ssh.ClientConfig{User: user, Auth: PassWd, HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	},
	}
	return ssh.Dial("tcp", ip_port, &Conf)
}
func dailPublic(user, publicKey, ipAdress string) (*ssh.Client, error) {
	if strings.HasSuffix(publicKey, ".pub") {
		publicKey = strings.TrimSuffix(publicKey, ".pub")
	}
	signer, err := readPrivateKey(publicKey)
	if err != nil {
		fmt.Println(err)
	}

	Conf := ssh.ClientConfig{
		User: user,
		Auth:  []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	return ssh.Dial("tcp", ipAdress, &Conf)
}
func readPrivateKey(path string) (ssh.Signer, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return ssh.ParsePrivateKey(b)
}
func getWeekDay() []string {
	now := time.Now()
	weekDay := []string{}
	startDate := now.Add(-time.Duration(now.Weekday()-1+7) * time.Hour * 24)
	for i:=0;i<7;i++{
		weekDay = append(weekDay,startDate.Add(time.Duration(i) * time.Hour * 24).Format("20060102"))
	}
	return weekDay
}
func writeExcel(file *xlsx.File, data map[string]map[string]float64,sortKey []string) (*xlsx.File, error) {
	var row *xlsx.Row
	if len(data) > 0 {
		sheet, err := file.AddSheet("Sheet1")
		if err != nil {
			return file, err
		}
		row = sheet.AddRow()
		row.AddCell().Value = "时间"
		row.AddCell().Value = "消费金额"
		row.AddCell().Value = "消费笔数"
		row.AddCell().Value = "积分数量"

		for _, key := range sortKey {
			if len(data[key])>0{
				stat := data[key]
				row = sheet.AddRow()
				row.AddCell().Value = key
				row.AddCell().Value = fmt.Sprintf("%.2f",stat["amount"])
				row.AddCell().Value = fmt.Sprintf("%v",stat["count"])
				row.AddCell().Value = fmt.Sprintf("%.2f",stat["score"])
			}
		}
	}

	return file, nil
}
func runCmd(Client *ssh.Client,week []string) (ret map[string]string) {
	ret = make(map[string]string,0)
	for _,v := range week{
		cmd := fmt.Sprintf("grep 'ConsumeScore req:' /opt/angrycard/logs/angrycard.log.%s",v)

		session, err := Client.NewSession()
		defer session.Close()
		if err != nil {
			fmt.Println("创建Session失败:", err)
			return
		}
		var stdOut, stdErr bytes.Buffer
		session.Stdout = &stdOut
		session.Stderr = &stdErr
		err = session.Run(cmd)
		if err !=nil && strings.Contains(err.Error(),"Process exited with status 1"){
			//fmt.Printf("%v is empty\n",v)
			ret[v] = ""
			continue
		}else if err != nil {
			//fmt.Printf("session run cmd err:%s",err)
			return
		}else{
			ss := strings.Replace(stdOut.String(), "\n", "", -1)
			ret[v] = ss
		}
	}
	return ret
}
func summaryAll(ret map[string]string)(summary map[string]map[string]float64){
	var score,count,pay float64 = 0,0,0
	summary = make(map[string]map[string]float64,0)
	var sli85 = make(map[string]string,0)
	var sli79 = make(map[string]string,0)
	split,str85,str79 := "func=mallcoo.(*api).ConsumeScore","consumeScore.go:85","consumeScore.go:79"
	findScore,findPay,findSuc := "\"Score\":","PayAmount:","\"Message\":\"成功\""
	for k,v := range ret {
		if v != "" {
			sli := strings.Split(v, split)
			p := len(sli)
			//反转数组
			for i, j := 0, p-1; i < j; i, j = i+1, j-1 {
				sli[i], sli[j] = sli[j], sli[i]
			}

			i:=0
			count,score = 0,0
			for{
				if i>p-1{
					break
				}

				if index85 := strings.Index(sli[i], str85);index85 != -1 {
					if indexSuc := strings.Index(sli[i], findSuc);indexSuc != -1 {
						val := subStrValue(sli[i], findScore, ",")
						temp,err := strconv.ParseFloat(val,10)
						if err != nil{
							fmt.Printf("85 ParseFloat err:%s\n",err)
							return
						}
						sli85[sli[i][39:39+7]] = "1"
						count++
						score += temp
					}
				}
				if index79 := strings.Index(sli[i], str79);index79 != -1{
					sli79[sli[i][39:39+7]] = sli[i]
				}
				i++
			}
			pay = 0
			for k1,_:= range sli85 {
				if len(sli79[k1])!=0{
					val := subStrValue(sli79[k1], findPay, " ")
					temp,err := strconv.ParseFloat(val,10)
					if err != nil{
						fmt.Printf("79 ParseFloat err:%s\n",err)
						return
					}
					pay += temp
				}
			}
			summary[k] = map[string]float64{"score":score,"amount":pay,"count":count}
		}else{
			summary[k] = map[string]float64{"score":0,"amount":0,"count":0}
		}
	}

	return summary
}
func sendEmail(file string) error {
	date := time.Now().Format("2006-01-02")
	emailbody := `
<html>
<body>
<p>Dear,<p>

<blockquote>
<p>附件是配劵页面%s活动信息内部报表，请查收。<p>
<p>内部文件，注意保密，严禁外传。<p>
<br>
<p>如有问题请及时联系。<p>
<blockquote>
</body>
</html>`
	em := email.NewEmail()

	em.To = []string{"carey.li@cardinfolink.com"}
	em.From = USER
	em.Subject = fmt.Sprintf("口碑平台配劵页面%s数据报表",date)
	//em.Cc = stringArray
	em.HTML = []byte(fmt.Sprintf(emailbody,date)) // Content-Type: text/html

	fileR,err := os.Open(file)
	em.Attach(fileR, "口碑平台配劵页面数据报表.xlsx","")

	err = em.Send(ADDR, author)
	if err != nil {
		fmt.Printf("TO %s, send email fail: %s", em.To, err)
	}

	return err
}