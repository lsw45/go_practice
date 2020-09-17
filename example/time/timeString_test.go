package time

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestTimeToString(t *testing.T) {
	ActiveTime := time.Now()
	InActiveTime := ActiveTime.Format("2006-01-02 15:04:05")
	fmt.Printf("%s\n%v\n", ActiveTime.String(), InActiveTime)

	str := []string{ActiveTime.String(), InActiveTime, "33333", "78787878"}
	jsonStr, _ := json.Marshal(str)
	t.Log(string(jsonStr))
}
