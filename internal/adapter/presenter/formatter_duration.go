package presenter

import (
	"fmt"
	"time"
)

func ToDurationString(d time.Duration) string {
	hour := uint(d.Hours())
	min := uint(d.Minutes()) % 60
	sec := uint(d.Seconds()) % 60
	return fmt.Sprintf("%d:%02d:%02d", hour, min, sec)
}
