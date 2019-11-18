package main

import (
	"fmt"
)

func main() {
	//多维map的声明与实现方法
	//方法1 初始化一个空的多维映射
	mainMapA := map[string]map[string]string{}
	subMapA := map[string]string{"A_Key_1": "A_SubValue_1", "A_Key_2": "A_SubValue_2"}
	mainMapA["MapA"] = subMapA
	fmt.Println("MultityMapA")
	for keyA, valA := range mainMapA {
		for subKeyA, subValA := range valA {
			fmt.Printf("mapName=%s	Key=%s	Value=%s\n", keyA, subKeyA, subValA)
		}
	}

	//方法2 使用make声明一个多维映射(等同一般声明)
	//var mainMap map[string]map[string]string
	mainMapB := make(map[string]map[string]string)
	//内部容器必须再次初始化才能使用
	subMapB := make(map[string]string)
	subMapB["B_Key_1"] = "B_SubValue_1"
	subMapB["B_Key_2"] = "B_SubValue_2"
	mainMapB["MapB"] = subMapB
	fmt.Println("\nMultityMapB")

	for keyB, valB := range mainMapB {
		for subKeyB, subValB := range valB {
			fmt.Printf("mapName=%s	Key=%s	Value=%s\n", keyB, subKeyB, subValB)
		}
	}

	/* 方法3 使用interface{}初始化一个一维映射
	 * 关键点：interface{} 可以代表任意类型
	 * 原理知识点:interface{} 就是一个空接口，所有类型都实现了这个接口，所以它可以代表所有类型
	 */
	//mainMapC := make(map[string]interface{})
	mainMapC := map[string]interface{}{}
	subMapC := make(map[string]string)
	subMapC["C_Key_1"] = "C_SubValue_1"
	subMapC["C_Key_2"] = "C_SubValue_2"
	mainMapC["MapC"] = subMapC
	fmt.Println("\nMultityMapC")
	for keyC, valC := range mainMapC {
		//此处必须实例化接口类型，即*.(map[string]string)
		//subMap := valC.(map[string]string)
		for subKeyC, subValC := range valC.(map[string]string) {
			fmt.Printf("mapName=%s	Key=%s	Value=%s\n", keyC, subKeyC, subValC)
		}
	}
}
