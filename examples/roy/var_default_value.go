// Default values of primitive types
package main

import "fmt"

func main() {
	var i int
	fmt.Println("default int is: ", i)

	var s string
	fmt.Println("default string is: ", s)

	var f float64
	fmt.Println("default float64 is: ", f)

	var arrInt [3]int
	fmt.Println("default int array is: ", arrInt)

	// var u unit64
	// fmt.Println("default unit64 is: ", u)

	// var c complext64
	// fmt.Println("default complext64 is: ", c)
}
