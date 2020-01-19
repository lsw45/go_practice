package polymorphism

import (
	"fmt"
	"testing"
)

type annimal interface {
	eat()
	sleep()
	run()
}

type cat interface {
	annimal
	Climb()
}
type dog interface {
	annimal
}

type HelloKitty struct {
	cat
}
type husky struct {
	dog
}

func (h HelloKitty) eat() {
	fmt.Println("eat cake!!")
}

func (h husky) eat() {
	fmt.Println("eat bone!!")
}

func TestAnimal(t *testing.T) {
	// var a annimal
	// a = HelloKitty{}
	// test(a)
	// a = husky{}
	// test(a)

	var a annimal
	a = HelloKitty{}
	var b annimal
	b = husky{}

	var animals [2]annimal = [...]annimal{a, b} //新建接口数组，放入结构体元素

	for _, v := range animals {
		if data, ok := v.(husky); ok { // 对象.(指定的类型) 判断改对象是否是指定的类型
			data.eat()
			fmt.Println("this is wangcai : ")
		}
		if data, ok := v.(HelloKitty); ok {
			data.eat()
			fmt.Println("this is HelloKitty : ")
		}
	}
}
