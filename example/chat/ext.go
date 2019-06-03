package main

import (
	"context"
	"fmt"
	"time"
)

func main() { // 传递带超时的上下文告诉阻塞函数// 超时过后应该放弃它的工作。
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // 打印"context deadline exceeded"
	}
}
