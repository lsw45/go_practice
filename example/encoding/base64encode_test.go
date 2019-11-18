package encoding

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestEncoding(t *testing.T) {
	str := "eyJjYXJkQ29kZSI6IjIwMTkxMDExMDkwMTMxMDA2NyIsImNoYW5uZWwiOiJBTFAiLCJmdW5jIjoibmV3Q2FyZCJ9"
	extInfoJsonStr, _ := base64.StdEncoding.DecodeString(str)
	fmt.Printf(string(extInfoJsonStr))
}
