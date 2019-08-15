package main

import (
	"fmt"
)

func main() {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
	main019()
}

func DeferFunc1(i int) (t int) {
	t = i

	defer func() {
		t += 3
	}()

	return t
}

func DeferFunc2(i int) int {
	t := i

	defer func() {
		t += 3
	}()

	return t
}

func DeferFunc3(i int) (t int) {

	defer func() {
		t += i
	}()

	return 2
}

func main019() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		name string
		age  int
	}{name: "qq", age: 11}

	/*
		比较结构体的相等，比较属性列表的顺序与值，
		调换了属性列表的顺序后，就不再视为相同的Type，就不能再比较了
	*/
	if sn1 == sn2 {
		fmt.Println("sn1== sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	/*属性中包含不能做==运算的类型时，结构体也无法进行==运算*/
	if sm1 == sm2 {
		fmt.Println("sm1== sm2")
	}
}
