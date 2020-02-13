package encoding

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"testing"
)

func TestEncoding(t *testing.T) {
	str := "eyJjYXJkQ29kZSI6IjIwMTkxMDExMDkwMTMxMDA2NyIsImNoYW5uZWwiOiJBTFAiLCJmdW5jIjoibmV3Q2FyZCJ9"
	extInfoJsonStr, _ := base64.StdEncoding.DecodeString(str)
	fmt.Printf(string(extInfoJsonStr))
}

func TestStrconv(t *testing.T) {
	int64_num := float32(6)
	int64_num2 := float32(9)
	a := int64_num / int64_num2
	f, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", a), 32)
	f1 := fmt.Sprintf("%.2f", a)
	fmt.Println(f)
	fmt.Println(f1)
}
