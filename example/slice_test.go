package main

import "fmt"
import "testing"

func main() {

	input := [...]int{1, 2, 3, 4, 6}            // 数组
	fmt.Printf("%T", input)                     // [5]int
	a := input[:]                               // 切片:
	fmt.Printf("%T,%d,%d\n", a, len(a), cap(a)) // []int,5,5

	// 创建普通切片
	var v []int = make([]int, 10)             // 等同于v := make([]int, 10),切片v现在是对一个新的有10个整数的数组的引用
	fmt.Printf("%v,%d,%d", v, len(v), cap(v)) // 输出：[0 0 0 0 0 0 0 0 0 0],10,10

	// 创建指向切片的指针
	var p *[]int = new([]int)
	fmt.Println(p) // 输出：&[]
	// fmt.Printf("%v,%d,%d", p, len(p), cap(p)) //报错; invalid argument p (type *[]int) for len
	fmt.Printf("%v,%d,%d", *p, len(*p), cap(*p)) // &[],0,0

	*p = make([]int, 10, 10)
	fmt.Println(p)       // 输出：&[0 0 0 0 0 0 0 0 0 0]
	fmt.Println((*p)[2]) // 输出： 0

	// *p := make([]int, 10, 10)//这么写会报错
	v = append(v, a...)                       //切片相加的写法
	fmt.Printf("%v,%d,%d", v, len(v), cap(v)) //[0 0 0 0 0 0 0 0 0 0 1 2 3 4 6],15,20

}

func TestSlice(t *testing.T) {

	var m map[string]int

	key := "two"
	elem, ok := m["two"]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n", key, elem, ok)

	fmt.Printf("The length of nil map: %d\n", len(m))

	fmt.Printf("Delete the key-element pair by key %q...\n", key)
	delete(m, key)

	elem = 2
	fmt.Println("Add a key-element pair to a nil map...")
	m["two"] = elem // 这里会引发panic。

}
