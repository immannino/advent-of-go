package util

import "fmt"

func Print(f func() string) {
	fmt.Println(f())
}

func PrintFunc(f func() string) func() {
	return func() {
		fmt.Println(f())
	}
}

func PrettyPrint(f func() string) {
	fmt.Printf("Pretty Print: %s\n", f())
}

func PrefixPrint(p string, f func() string) {
	fmt.Printf("%s %s", p, f())
}
