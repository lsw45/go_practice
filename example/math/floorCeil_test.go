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
	fmt.Println(float64(192) / float64(193))
	fmt.Println(int64(192) / int64(193))
	fmt.Println(1 / 193)

	lc := make([]int, 0)
	//lc := []int{0, 1, 2, 3, 4}
	for _, v := range lc {
		fmt.Println(v)
	}
}
