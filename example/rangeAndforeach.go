package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	m1 := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		stu.Age = stu.Age + 10
	}
	fmt.Println(stus)
	fmt.Printf("%+v", m["li"])
	fmt.Println()
	for i := 0; i < len(stus); i++ {
		m1[stus[i].Name] = &stus[i]
	}
	fmt.Println(m1)
	fmt.Printf("%+v", m["li"])
}
