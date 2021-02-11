package year2015

import (
	"fmt"
	"testing"
)

// TestOvo testing to help myself
func TestToggleLights(t *testing.T) {
	lights := make([]uint8, 1000000)
	lights1 := make([]uint8, 1000000)
	lights2 := make([]uint8, 1000000)

	var tests = []struct {
		lights  *[]uint8
		vals    *Coords
		action  string
		rowSize int
		count   int
	}{
		{
			lights: &lights,
			vals: &Coords{
				aStart: 0,
				aEnd:   0,
				bStart: 999,
				bEnd:   999,
			},
			rowSize: 1000,
			action:  "TURNON",
			count:   1000000,
		},
		{
			lights: &lights1,
			vals: &Coords{
				aStart: 0,
				aEnd:   0,
				bStart: 999,
				bEnd:   0,
			},
			rowSize: 1000,
			action:  "TOGGLE",
			count:   1000,
		},
		{
			lights: &lights2,
			vals: &Coords{
				aStart: 499,
				aEnd:   499,
				bStart: 500,
				bEnd:   500,
			},
			rowSize: 1000,
			action:  "TURNOFF",
			count:   0,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.action)

		t.Run(testname, func(t *testing.T) {
			toggleLights(tt.lights, tt.vals, tt.action, tt.rowSize)

			actual := 0

			for _, c := range *tt.lights {
				if c == 1 {
					actual++
				}
			}

			fmt.Println(actual, tt.count)
			if actual != tt.count {
				t.Errorf("got %d, expected %d", actual, tt.count)
			}
		})
	}
}
