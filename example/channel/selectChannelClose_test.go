package channel

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 需求功能：for循环去channel，当channel关闭，for马上退出
func TestChannelClose(t *testing.T) {
	channel := make(chan string)
	timerand := time.Now().Unix()
	rand.Seed(timerand) //If Seed is not called, the generator behaves as if seeded by Seed(1).

	go func() {
		cnt := rand.Intn(55)
		fmt.Println("message cnt :", cnt)
		for i := 0; i < cnt; i++ {
			channel <- fmt.Sprintf("message-%2d", i)
		}
		close(channel) //关闭Channel
	}()

	var more bool = true
	var msg string
	for more {
		select {
		//channel会返回两个值，一个是内容，一个是还有没有内容
		case msg, more = <-channel:
			if more {
				fmt.Println(msg)
			} else {
				fmt.Println("channel closed!")
			}
		}
	}
}
