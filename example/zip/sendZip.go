package zip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

type Files struct {
	Buf  *bytes.Buffer
	Name string
}

func TestSendZipMail(fs []Files, title, body string) {
	yesterday := time.Now().AddDate(0, 0, -1)
	date := yesterday.Format("20060102")
	e := &email.Email{
		To:    "carey.li@cardinfolink.com",
		Cc:    "carey.li@cardinfolink.com",
		Title: title,
		Body:  fmt.Sprintf(body, date),
	}

	var err error
	tmpFile, _ := os.Create("./tmpFile.zip")
	defer tmpFile.Close()

	zw := zip.NewWriter(tmpFile)
	for _, file := range fs {
		fw, err := zw.Create(file.Name)
		if err != nil {
			logrus.Errorf("新增压缩文件失败：%s", err)
			return
		}

		_, err = fw.Write(file.Buf.Bytes())
		if err != nil {
			logrus.Errorf("写入压缩文件失败：%s", err)
			return
		}
	}
	_ = zw.Close()

	_, _ = tmpFile.Seek(0, 0) //设置偏移量到文件头部
	e.Attach(tmpFile, "积分报表汇总.zip", "")

	// 如果失败 就重试5次
	var retryCount = 0
	for {
		if retryCount >= 5 {
			log.Printf("send report fail, err:%s", err)
			break
		}
		err = e.Send()
		if err != nil {
			retryCount++
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
}
