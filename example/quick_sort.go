package main

import (
	"fmt"
)

// 1、快速排序关键过程是对数组进行划分，划分过程需要选择一个主元素(第一个或最后一个)作为参照，
// 2、所有小于主元素的都移到主元素左边，大的移到右边
// 3、对主元素左右两边的两个子集，不断重复第一和第二步，直到所有子集只剩下一个元素为止

func main() {
	array := []int{3, 1, 2, 5, 4}
	fmt.Println(array)

	// 单线程实现快速排序
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)

	// 多线程实现快速排序
	ch := make(chan int)
	go quickSortChan(array, ch)
	for value := range ch {
		fmt.Println(value)
	}
}

// 实现第二步
func partition(array []int, left, right int) int {
	baseNum := array[left] //选择第一个为主元素
	//按递增排序：即小于主元素的放在左边，大的放右边
	// 当left==right时，退出循环
	for left < right {
		//遍历右边直到找到小于主元素的值
		for array[right] >= baseNum && right > left {
			fmt.Printf("right:%d\n", right)
			right--
		}
		//将较小值放在原主元素的位置上
		array[left] = array[right]
		fmt.Printf("right:")
		fmt.Println(array)

		//遍历左边直到大于主元素值
		for array[left] <= baseNum && right > left {
			fmt.Printf("left:%d\n", left)
			left++
		}
		//将较大值放在主元素右边
		array[right] = array[left]
		fmt.Printf("left:")
		fmt.Println(array)
	}

	array[right] = baseNum //确定主元素的位置
	fmt.Printf("确定主元素的位置")
	fmt.Println(array)
	return right //返回主元素的位置
}

// 借助递归实现第三步
func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}
	index := partition(array, left, right)
	fmt.Printf("index------%d\n", index)
	quickSort(array, left, index-1)
	quickSort(array, index+1, right)
}

func quickSortChan(array []int, ch chan int) {
	if len(array) == 1 {
		ch <- array[0]
		close(ch)
		return
	}
	if len(array) == 0 {
		close(ch)
		return
	}
	small := make([]int, 0)
	big := make([]int, 0)
	left := array[0]
	array = array[1:]
	for _, num := range array {
		switch {
		case num <= left:
			small = append(small, num)
		case num > left:
			big = append(big, num)
		}
	}
	left_ch := make(chan int, len(small))
	right_ch := make(chan int, len(big))

	go quickSortChan(small, left_ch)
	go quickSortChan(big, right_ch)

	//合并数据
	for i := range left_ch {
		ch <- i
	}
	ch <- left
	for i := range right_ch {
		ch <- i
	}
	close(ch)
	return
}

/*
[3 1 2 5 4]
right:4
right:3
right:[2 1 2 5 4]
left:0
left:1
left:[2 1 2 5 4]
确定主元素的位置[2 1 3 5 4]
index------2
right:[1 1 3 5 4]
left:0
left:[1 1 3 5 4]
确定主元素的位置[1 2 3 5 4]
index------1
right:[1 2 3 4 4]
left:3
left:[1 2 3 4 4]
确定主元素的位置[1 2 3 4 5]
index------4
[1 2 3 4 5]
1
2
3
4
5
*/
