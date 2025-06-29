package greeting2

import (
	"time"
)

func Do() string {
	hour := time.Now().Hour()

	switch {
	case hour >= 4 && hour < 10:
		return "おはよう"
	case hour < 17:
		return "こんにちは"
	case hour < 20:
		return "こんばんは"
	default:
		return "おやすみ"
	}
}
