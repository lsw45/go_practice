package operator

import (
	"fmt"
	"testing"
)

//
// <<：左移n位就是乘以2的n次方
// >>：右移n位就是除以2的n次方
// &与。 |或
// 异或运算符^：二进位相异时，结果为1
func TestOr(t *testing.T) {
	fmt.Println(^0)           //-1
	fmt.Println(^uint(0))     //18446744073709551615
	fmt.Println(uint(0) >> 1) // 0

	maxInt := int(^uint(0) >> 1) //9223372036854775807
	fmt.Println(maxInt)

	fmt.Println(^uint(0) / uint(maxInt)) // 2
}

// ^：对于任何数x，都有x^x=0，x^0=x
func TestOr1(t *testing.T) {
	fmt.Printf("%b\n", 2)  //10
	fmt.Printf("%b\n", ^2) //-11

	fmt.Printf("%b\n", 8)  //1000
	fmt.Printf("%b\n", ^8) //-1001

	fmt.Printf("%b\n", -8)  //-1000
	fmt.Printf("%b\n", ^-8) //111

	fmt.Printf("%b\n", 89)  //1011001
	fmt.Printf("%b\n", ^89) //-1011010

	fmt.Printf("%b\n", uint(2))  //10
	fmt.Printf("%b\n", ^uint(2)) // 1111111111111111111111111111111111111111111111111111111111111101

	fmt.Printf("%b\n", uint(9))  //1001
	fmt.Printf("%b\n", ^uint(9)) //1111111111111111111111111111111111111111111111111111111111110110

}

/*1-1000放在含有1001个元素的数组中，只有唯一的一个元素值重复，其它均只出现
一次。每个数组元素只能访问一次，设计一个算法，将它找出来；不用辅助存储空
间，能否设计一个算法实现？
答案：
将所有元素全部异或，得到的结果与1^2^3^...^1000的结果进行异或，得到的结果就是重复数。*/
