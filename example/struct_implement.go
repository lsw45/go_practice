package main

import "fmt"

type People3 interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {}

func live() People3 {
	var stu *Student
	fmt.Println(stu)
	return stu
}

func main() {
	people3 := live()
	fmt.Println(people3)

	var p *int
	fmt.Println(p)
	if p == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}

	if people3 == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
	funs := test()
	for _, f := range funs {
		f()
	}
}

func test() []func() {
	var funs []func()
	for i := 0; i < 10; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}
