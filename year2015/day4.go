package year2015

import (
	"crypto/md5"
	"fmt"
	"io"
)

var data4 string = "yzbqklnj"

// Welp time to learn how MD5 hash works.
func Day4() {
	h := md5.New()
	// secret := []byte(data4)

	io.WriteString(h, "00000")

	// fmt.Printf("%x", h.Sum(secret))

	// fmt.Printf("Day 4: { 1: %d, 2: %d }\n", h.Sum(nil), 0)
	fmt.Printf("Day 4: Learn MD5 hash\n")
}
