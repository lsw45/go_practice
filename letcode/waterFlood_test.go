package letcode

import (
	"fmt"
	"testing"
)

func TestWaterFlood(t *testing.T) {
	rains := []int{1, 20, 0, 1, 2}
	n := len(rains)
	res := make([]int, n)
	zeros := make([]int, n)
	map1 := make(map[int]int)

	for i := 0; i < n; i++ {
		if rains[i] == 0 {
			zeros = append(zeros, i)
		} else {
			res[i] = -1
			pool := rains[i]
			if lastFull, flag := map1[pool]; flag {
				id := lower_bound(zeros, lastFull)
				zeroIdx := 0
				if id > len(zeros) {
					zeroIdx = -1
				} else {
					zeroIdx = zeros[id]
				}
				if zeroIdx > lastFull {
					res[zeroIdx] = pool
					zeros = append(zeros[:zeroIdx], zeros[zeroIdx+1:]...)
					map1[pool] = i
				} else {
					fmt.Println([]int{0})
				}
			} else {
				map1[pool] = i
			}
		}
	}
}

/**二分搜索**/
func lower_bound(A []int, target int) int {
	lo := 0
	hi := len(A)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if A[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
