// Control structures - Go for loop, break, continue, range
package main

import (
	"fmt"
)

func first() {
	// 'for' statement like C style syntax
	for i := 0; i < 5; i += 1 {
		fmt.Println("value of i is now ", i)
	}
}

func second() {
	// infinite looping
	// DO NOT EXECUTE THIS FUNCTION
	for i := 0; ; i++ {
		fmt.Println("value of i is now ", i)
	}
}

func third() {
	// infinite looping
	// DO NOT EXECUTE THIS FUNCTION
	for i := 1; i < 3; {
		fmt.Println("value of i is now ", i)
	}
}

func forth() {
	// possible string comparison for condition
	s := nil
	for s != "aaaaa" {
		fmt.Println("value of s is ", s)
		s = s + "a"
	}
}

func main() {
	first()
	forth()
}
