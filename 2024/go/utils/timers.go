package utils

import (
	"time"
)

func WithTimer(partFunc func() (int, error)) (int, time.Duration, error) {
	T1 := time.Now()
	result, err := partFunc()
	if err != nil {
		return -1, -1, err
	}
	elapsed := -time.Until(T1)

	return result, elapsed, nil
}
