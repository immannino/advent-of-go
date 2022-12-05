package util

import (
	"fmt"
	"time"
)

func Time(f func() string) string {
	s := time.Now()

	o := f()

	e := time.Since(s)

	return fmt.Sprintf("%s - %v", o, e)
}
