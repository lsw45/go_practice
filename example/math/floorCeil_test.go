package math

import (
	"fmt"
	"math"
	"testing"
)

func TestFloor(t *testing.T) {
	math.Floor(3 / 4)
	//math.Floor(int64(4) / int64(4))
	//math.Floor(int(4) / int(4))
	fmt.Printf("float64(192) / float64(193):%v\n", float64(192)/float64(193))
	fmt.Printf("int64(192) / int64(193):%v\n", int64(192)/int64(193))
	fmt.Printf("1 / 193:%v\n", 1/193)

	lc := make([]int, 0)
	//lc := []int{0, 1, 2, 3, 4}
	for _, _ = range lc {
		fmt.Println(lc)
	}

	fmt.Printf("float64(1/100):%v\n", (float64(1)/float64(10))*float64(100))
	i := int64(math.Floor(0.1))
	fmt.Println(i)
}
