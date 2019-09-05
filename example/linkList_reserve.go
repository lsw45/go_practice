package main

// 反转链表

import (
	"fmt"
)

type Linked struct {
	value int
	next  *Linked
}

// 生成头结点，
func New() *Linked {
	return &Linked{}
	//或者 return &Linked{0,nil}
}

// 在某位置插入节点
func (head *Linked) Insert(i, n int) bool {
	p := head
	j := 0
	for nil != p && j < i {
		p = p.next
		j++
	}

	if nil == p || j > i {
		fmt.Println("")
		return false
	}

	s := &Linked{value: n}
	s.next = p.next
	p.next = s

	return true
}

func (head *Linked) Traverse() {
	point := head
	for nil != point {
		fmt.Println(point.value)
		point = point.next
	}
	fmt.Println("--------done----------")
}

func main() {
	linkedList := New()
	linkedList.Insert(0, 9)
	linkedList.Traverse()

	fmt.Println("xxx")
}
