package letcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 4, 9, 56, 90}
	target := 8

	numbersSize := len(numbers)
	returnSize := []int{0, 0}

	i, j := 0, 0
	if (numbers[0] > target) || (numbers[0]+numbers[1] > target) || (numbers[numbersSize-2]+numbers[numbersSize-1] < target) {
		t.Logf("%+v", returnSize)
		return
	} else if numbers[0]+numbers[1] == target {
		returnSize = []int{1, 2}
		t.Logf("%+v", returnSize)
		return
	} else if numbers[numbersSize-2]+numbers[numbersSize-1] == target {
		returnSize = []int{numbersSize - 1, numbersSize}
		t.Logf("%+v", returnSize)
		return
	}

	tag := true

	for i = 0; i < numbersSize; i++ {
		for j = numbersSize - 1; j > i; j-- {
			if numbers[i]+numbers[j] == target {
				returnSize = []int{i + 1, j + 1}
				tag = false
			}
		}
		if !tag {
			t.Logf("%+v", returnSize)
			return
		}
	}
	t.Logf("%+v", returnSize)
}
