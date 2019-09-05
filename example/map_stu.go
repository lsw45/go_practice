package main

import (
	"fmt"
	"runtime"
)

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)

	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus {
		// fmt.Println(stu)
		m[stu.Name] = &stu
	}

	fmt.Println(m)
	fmt.Println(runtime.NumCPU()) //4
}
