package date

import (
	"time"
)

func TimeToMarketClose() time.Duration {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	marketClose := time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, loc)
	result := marketClose.Sub(now)
	return result
}
