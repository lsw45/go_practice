package main

import (
	"fmt"

	_ "time"
)

type int33 = int
type inth int

func main() {
	var li int33 = 23
	var rou inth = 22
	alias(li)  // int 23
	alias(rou) // 报错： cannot use rou (type inth) as type int in argument to alias
}

func alias(pa int) {
	fmt.Printf("%T\n", pa)
	fmt.Printf("%v\n", pa)
}
