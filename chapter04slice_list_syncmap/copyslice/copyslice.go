package main

import "fmt"

func main() {
	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	copyData := make([]int, elementCount)

	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	// 引用切片数据
	refData := srcData
	// 将数据拷贝到新的切片空间中
	n := copy(copyData, srcData)
	fmt.Println(n) //拷贝数量：1000

	// 修改原数据的第一个元素
	srcData[0] = 99
	fmt.Println(refData[0])                            //99
	fmt.Println(copyData[0], copyData[elementCount-1]) //0 999

	// 拷贝原数据从4到6（不包含）
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i]) //0 1 2 3 4
	}
	fmt.Printf("\n")
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i]) //4 5 2 3 4
	}

	fmt.Printf("\n地址：%p %p %p\n", srcData, refData, copyData)                //地址不同：0xc000086000 0xc000086000 0xc000088000
	fmt.Printf("\n原切片的长度和容量：%d %d\n", len(srcData), cap(srcData))            //原切片的长度和容量：1000 1000
	fmt.Printf("拷贝srcData[4:6]的切片长度和容量：%d %d", len(copyData), cap(copyData)) //拷贝srcData[4:6]的切片长度和容量：1000 1000

	refCap()
}

func refCap() {
	fmt.Println("\n     len cap   address")

	nums := []int{1, 2, 3, 4, 5}
	fmt.Print("nums---", len(nums), cap(nums)) //nums---5 5
	fmt.Printf("    %p\n", nums)               //0xc4200181e0

	numa := nums[:3:4]
	fmt.Print("numa---", len(numa), cap(numa)) //numa---3 4
	fmt.Printf("    %p\n", numa)               //0xc4200181e0 地址一样

	numb := make([]int, 3)
	copy(numb, nums[:2:4])
	fmt.Print("numb---", len(numb), cap(numb)) //numb---3 3
	fmt.Printf("    %p\n", numb)               //0xc000050140 地址不一样

	nums[0] = 55
	fmt.Println(nums, numa, numb) //[55 2 3 4 5] [55 2 3] [1 2 0]
}
