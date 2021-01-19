package year2015

import (
	"fmt"
	"testing"
)

// TestOvo testing to help myself
func TestOvo(t *testing.T) {
	var tests = []struct {
		word string
		want bool
	}{
		{"aaa", true},
		{"abc", false},
		{"ovo", true},
		{"ovov", true},
		{"pwpt", true},
		{"aa", false},
		{"a", false},
		{"1234", false},
		{"alksdjfovosdfds", true},
		{"yktkmkokmfslgkml", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.word)

		t.Run(testname, func(t *testing.T) {
			actual := ovo(tt.word)

			if actual != tt.want {
				t.Errorf("got %t, expected %t", actual, tt.want)
			}
		})
	}
}

func TestHasOverlappingPairs(t *testing.T) {
	var tests = []struct {
		word string
		want bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"xyxy", true},
		{"aabcdefgaa", true},
		{"uurcxstgmygtbstg", true},
		{"ieodomkazucvgmuy", false},
		{"tonyggggtony", true},
		{"aaa", false},
		{"yktkmkokmfslgkml", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.word)

		t.Run(testname, func(t *testing.T) {
			actual := hasOverlappingPairs(tt.word)

			if actual != tt.want {
				t.Errorf("got %t, expected %t", actual, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	var tests = []struct {
		word string
		want bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"suerykeptdsutidb", false},
		{"yktkmkokmfslgkml", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.word)

		t.Run(testname, func(t *testing.T) {
			actual := (hasOverlappingPairs(tt.word) && ovo(tt.word))

			if actual != tt.want {
				t.Errorf("got %t, expected %t", actual, tt.want)
			}
		})
	}
}
