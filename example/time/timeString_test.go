package time

import (
	"fmt"
	"testing"
	"time"
)

type Module struct {
	ActiveTime   *time.Time `json:"activeTime,omitempty" gorm:"not null" required:"true" description:"上架时间"`
	InActiveTime string     `json:"inactiveTime,omitempty" gorm:"column:inactive_time;not null" required:"true" description:"下架时间"`
	CreatedAt    *time.Time
}

func TestTimeToString(t *testing.T) {

	now := time.Now()
	later := now.AddDate(100, 0, 0)

	adv := Module{
		ActiveTime:   &now,
		InActiveTime: later.Format("2006-01-02 15:04:05"),
		CreatedAt:    &now,
	}
	fmt.Printf("%s\n%v\n%v\n", adv.ActiveTime.String(), adv.InActiveTime, adv.CreatedAt)
}
