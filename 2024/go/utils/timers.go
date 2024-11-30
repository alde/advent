package utils

import (
	"time"
)

func WithTimer(partFunc func() int) (int, time.Duration) {
	T1 := time.Now()
	result := partFunc()
	elapsed := -time.Until(T1)

	return result, elapsed
}
