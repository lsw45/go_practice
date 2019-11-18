package main

import "fmt"
import "testing"

// func TestNew(t *testing.T) {
func main() {
	channel := make(chan string)
	fmt.Printf("%T,%d,%d\n", channel, len(channel), cap(channel)) // chan string,0,0

	input := [...]int{1, 2, 3, 4, 6}                        // 数组
	fmt.Printf("%T,%d,%d\n", input, len(input), cap(input)) // [5]int,5,5
	// sortArray := [10]int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36} // 报错：cannot use []int literal (type []int) as type [10]int in assignment

	a := input[:]                               // 切片:
	fmt.Printf("%T,%d,%d\n", a, len(a), cap(a)) // []int,5,5

	// 创建普通切片
	var v []int = make([]int, 10)             // 等同于v := make([]int, 10),切片v现在是对一个新的有10个整数的数组的引用
	fmt.Printf("%v,%d,%d", v, len(v), cap(v)) // 输出：[0 0 0 0 0 0 0 0 0 0],10,10

	// 创建并初始化一个切片
	sortArray := []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}
	fmt.Printf("%T,%d,%d\n", sortArray, len(sortArray), cap(sortArray)) //[]int,10,10

	// 创建指向切片的指针
	var p *[]int = new([]int)
	fmt.Println(p) // 输出：&[]
	// fmt.Printf("%v,%d,%d", p, len(p), cap(p)) //报错; invalid argument p (type *[]int) for len
	fmt.Printf("%v,%d,%d", *p, len(*p), cap(*p)) // [],0,0

	t := make([]int, 0)
	fmt.Printf("%v,%d,%d", t, len(t), cap(t)) //[],0,0

	// *t := make([]int, 5, 10)//non-name *t on left side of :=，错误的声明方式
	*p = make([]int, 5, 10)
	fmt.Println(p)       // 输出：&[0 0 0 0 0]
	fmt.Println((*p)[2]) // 输出： 0
	// fmt.Println((*p)[7]) // panic: runtime error: index out of range

	v = append(v, a...)                       //切片相加的写法
	fmt.Printf("%v,%d,%d", v, len(v), cap(v)) //注意这里是10个0：[0 0 0 0 0 0 0 0 0 0 1 2 3 4 6],15,20

	s := map[string]int{} //等同于var s map[string]int;s = map[string]int{}
	s["xx"] = 3
	fmt.Printf("%v,%d", s, len(s)) //map[xx:3],1

	m := make(map[string]int)
	m["one"] = 1
	fmt.Printf("%v,%d", m, len(m)) //map[one:1],1
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
