package string_test

import (
	"fmt"
	_ "strconv"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	tracer := "\"死神来了,死神bye bye,xxx"
	comma := strings.Index(tracer, "\"")
	fmt.Println(comma)
	//comma的意思是从字符串tracer查找第一个逗号，然后返回他的位置，这里的每个中文是占3个字符，从0开始计算，那么逗号的位置就是12

	pos := strings.Index(tracer, ",")
	fmt.Println(pos)
	//tracer[comma:]这个是的意思截取字符串tracer，从12开始，包括12

	fmt.Println(strings.Trim(tracer[comma:pos], "\"")) //可以直接+1：tracer[comma+1:pos]
	//,死神bye bye

	//整段的代码的意思是从tracer[comma:]这个字符串中查找“死神”这个字符串，第0位是逗号，第一位开始就是“死神”了，所以这里pos是1

	// fmt.Println(tracer[comma+pos+3:])

	fmt.Println(len("] [source=mallcoo/consumeScore.go:85] ["))
	fmt.Println(len("3138326"))
}
