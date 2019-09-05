package main

import "time"
import "fmt"

func main() {
	//创建两个channel - c1 c2
	c1 := make(chan string)
	c2 := make(chan string)
	t1 := time.Now()

	//创建两个goruntine来分别向这两个channel发送数据
	go func() {
		fmt.Printf("hello")
		c1 <- "Hello"
	}()
	go func() {
		fmt.Printf("world")
		c2 <- "World"
	}()

	time.Sleep(time.Second * 5) //停止程序等待上面两个协程启动
	//使用select来侦听两个channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// 或者default，会导致无阻塞
	/*for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default: //default会导致无阻塞
			fmt.Println("nothing received!")
			time.Sleep(time.Second)
		}
	}*/

	// 也可以设置阻塞的Timeout
	timeout_cnt := 0
Loop:
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("msg1 received", msg1)
		case msg2 := <-c2:
			fmt.Println("msg2 received", msg2)
		case <-time.After(time.Second * 2):
			fmt.Println("Time Out")
			timeout_cnt++

			if timeout_cnt > 3 {
				break Loop
			}
		}
	}
	fmt.Println("time elapsed ", time.Since(t1))
}
