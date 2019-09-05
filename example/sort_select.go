package main

import (
	"fmt"
)

func main() {
	s := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	fmt.Println(s)
	// SelectionSort(s)
	SelectionSort(s, 10)
	fmt.Println(s)
}

/*//选择排序
func SelectionSort(s []int) {
	l := len(s) //以免每次循环判断都运算
	m := len(s) - 1
	for i := 0; i < m; i++ {
		k := i
		for j := i + 1; j < l; j++ {
			if s[j] < s[k] {
				k = j
			}
		}
		if k != i {
			//交换数组元素
			s[k], s[i] = s[i], s[k]
		}
	}
}*/

func SelectionSort(arr []int, n int) {
	k := 0
	for i := 0; i < n-1; i++ {
		k = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[k] {
				k = j
			}
		}
		if k != i {
			arr[k], arr[i] = arr[i], arr[k]
		}
	}
}
