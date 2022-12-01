package util

import "time"

func Time(f func()) time.Duration {
	s := time.Now()

	f()

	e := time.Since(s)

	return e
}
