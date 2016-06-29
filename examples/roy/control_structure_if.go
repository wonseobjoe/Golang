// Control structure - if and if-else
package main

import "fmt"

func first() {
	if true {
		fmt.Println("the true block is executed")
	}

	if false {
		fmt.Println("the false block won't be executed")
	}
}

func second() {
	a, b := 4, 5

	if a < b {
		fmt.Println(a, " is less than ", b)
	} else if a > b {
		fmt.Println(a, " is greater than ", b)
	} else {
		fmt.Println(a, " is equal to ", b)
	}
}

func main() {
	first()
	second()
}
