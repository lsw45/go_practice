package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {

	fmt.Println(runtime.NumCPU()) //4

	// append与切片扩容
	arr := [...]string{8: "Go", 2: "python", "java", "c", "c++", "php"}
	fmt.Printf("%+v\n", arr)

	si := arr[:5]
	fmt.Printf("%+v", si)
	fmt.Println("%d", len(si))

	si = append(si, "Ruby", "Erlang")
	fmt.Printf("%+v", si)
	fmt.Printf("%d\n", len(si))
	fmt.Printf("%+v\n", arr)
	/*
		[  python java c c++ php  Go]
		[  python java c]%d 5
		[  python java c Ruby Erlang]7
		[  python java c Ruby Erlang  Go]
	*/
	var numbers3 = [5]int{1, 2, 3, 4, 5}
	var slice1 = numbers3[1:3:4]                                //第三个参数：容量索引上界
	fmt.Printf("%+v,%d,%d\n", slice1, len(slice1), cap(slice1)) //[2 3],2,3
	slice1 = append(slice1, 6, 7)
	fmt.Printf("%+v,%d,%d\n", slice1, len(slice1), cap(slice1)) //[2 3 6 7],4,6
	fmt.Printf("%+v\n", numbers3)                               //[1 2 3 4 5]

	var pp = struct {
		Name string
		Age  uint8
	}{"Robert", 23}
	var puptr = uintptr(unsafe.Pointer(&pp))
	var nppName = puptr + unsafe.Offsetof(pp.Name)
	var nppAge = puptr + unsafe.Offsetof(pp.Age)
	fmt.Printf("%d\n", puptr)   //824634244632
	fmt.Printf("%d\n", nppName) //824634244632
	fmt.Printf("%d\n", nppAge)  //824634244648

	var name *string = (*string)(unsafe.Pointer(nppName))
	fmt.Printf("%v\n", *name) //Robert

	if uintptr(unsafe.Pointer(&pp))+unsafe.Offsetof(pp.Age) == uintptr(unsafe.Pointer(&pp.Age)) {
		fmt.Println(true)
	}
}
