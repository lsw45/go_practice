package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeToString(t *testing.T) {
	ActiveTime := time.Now()
	InActiveTime := ActiveTime.Format("2006-01-02 15:04:05")
	fmt.Printf("%s\n%v\n", ActiveTime.String(), InActiveTime)
}
