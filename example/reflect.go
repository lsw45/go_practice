package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num float64 = 1.2345
	fmt.Println("old value of pointer:", num)

	// 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
	pointer := reflect.ValueOf(&num)
	fmt.Println("pointer:", pointer)

	newValue := pointer.Elem()
	fmt.Println("newValue:", newValue)

	fmt.Println("type of pointer:", newValue.Type())
	fmt.Println("settability of pointer:", newValue.CanSet())

	// 重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num)

	////////////////////
	// 如果reflect.ValueOf的参数不是指针，会如何？
	pointer = reflect.ValueOf(num)
	//newValue = pointer.Elem() // 如果非指针，这里直接panic，“panic: reflect: call of reflect.Value.Elem on float64 Value”

	// 函数
	var jquery func(int) (error, int, float64)
	f := reflect.ValueOf(&jquery).Elem()
	fmt.Println("in of f:", f.Type().NumIn())   //函数的参数个数
	fmt.Println("out of f:", f.Type().NumOut()) //函数的返回值个数

	args := make([]reflect.Value, f.Type().NumOut())
	for i := 0; i < len(args); i++ {
		// args[i] = reflect.Zero(f.Type().Out(i)) //func Zero(type Type) Value：返回指定类型的零值
		args[i] = reflect.ValueOf(3)
	}

	fmt.Println("args list:", args)

	// var hlist []reflect.Value
	inArgs := make([]interface{}, 0, len(args))
	for _, v := range args {
		inArgs = append(inArgs, v.Interface())
	}
	fmt.Println("inArgs list:", inArgs)

	// realValue := value.Interface().(已知的类型)
	var hli float64 = 1.2345
	pointer = reflect.ValueOf(&hli)
	value := reflect.ValueOf(hli)
	// Go 对类型要求非常严格，类型一定要完全符合，如下两个，一个是*float64，一个是float64，如果弄混，则会panic
	fmt.Println("pointer value:", pointer.Interface().(*float64))
	fmt.Println("value value:", value.Interface().(float64))
}
