package year2015

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

var data5 string

func init() {
	data5 = utils.GetData("year2015/day5.txt")
}

// Welp time to learn how MD5 hash works.
func Day5() {
	list := strings.Split(data5, "\n")
	count := 0
	count2 := 0

	for _, c := range list {
		// part 1
		if hasThreeVowel(c) && hasRepeatingChar(c) && doesNotHaveSpecificStrings(c) {
			count += 1
		}

		if hasOverlappingPairs(c) && ovo(c) {
			fmt.Printf("%s - 1: %t 2: %t \n", c, hasOverlappingPairs(c), ovo(c))
			count2 += 1
		}
	}

	// fmt.Printf("Day 4: { 1: %d, 2: %d }\n", h.Sum(nil), 0)
	fmt.Printf("Day 5: { 1: %d, 2: %d }\n", count, count2)
}

// ---- Part 1

// Checks for the presence of at least 3 vowels
// Supported: 'aeiou'
func hasThreeVowel(word string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	count := 0

	for _, c := range word {
		if contains(&vowels, c) {
			count += 1
		}
	}

	if count >= 3 {
		return true
	}

	return false
}

func contains(valid *[]rune, letter rune) bool {
	for _, c := range *valid {
		if letter == c {
			return true
		}
	}

	return false
}

func hasRepeatingChar(word string) bool {
	currentChar := rune(word[0])

	for i := 1; i < len(word); i++ {
		tempChar := rune(word[i])

		if tempChar == currentChar {
			return true
		} else {
			currentChar = tempChar
		}
	}

	return false
}

// Checks for the absence of specific strings:
// Strings: [ ab, cd, pq, or xy ]
func doesNotHaveSpecificStrings(word string) bool {
	current := rune(word[0])
	badVals := []string{"ab", "cd", "pq", "xy"}

	for i := 1; i < len(word); i++ {
		temp := rune(word[i])

		str := string([]rune{current, temp})

		if isBadString(&badVals, str) {
			return false
		} else {
			current = temp
		}
	}

	return true
}

func isBadString(bad *[]string, word string) bool {
	for _, c := range *bad {
		if c == word {
			return true
		}
	}

	return false
}

// -- Part 2

// Create a map that keeps count of pairs. Return true if any keys have count 2+
func hasOverlappingPairs(word string) bool {
	patterns := make(map[string]int)

	for i := 0; i < (len(word) - 1); i++ {
		substr := word[i:(i + 2)]
		if _, ok := patterns[substr]; ok {
			patterns[substr] += 1
		} else {
			patterns[substr] = 1
		}
	}

	for _, v := range patterns {
		if v >= 2 {
			return true
		}
	}

	return false
}

// Returns true if any 3 chars follow the OvO pattern
func ovo(word string) bool {
	for i := 0; i < (len(word) - 2); i++ {
		first := word[i]
		// middle := word[i+1]
		last := word[i+2]

		// fmt.Printf("%s%s%s\n", string(first), string(middle), string(last))

		if first == last {
			return true
		}
	}

	return false
}
