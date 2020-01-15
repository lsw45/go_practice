package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	// 尾部添加
	l.PushBack("canon")

	// 头部添加
	l.PushFront(67)

	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")

	// 在“fist”之后添加”high”
	l.InsertAfter("high", element)

	// 在“fist”之前添加”noon”
	l.InsertBefore("noon", element)

	// 删除
	l.Remove(element)

	for p := l.Front(); p != nil; p = p.Next() {
		fmt.Println("Number", p.Value)
	}
}
