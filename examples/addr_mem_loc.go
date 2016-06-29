// Address and memory location
package main

import "fmt"

func main() {
	i := 5
	fmt.Println("i is: ", i)
	fmt.Println("address of i is: ", &i)

	ptr := &i
	fmt.Println("address of pointer to i: ", ptr)
	fmt.Println("value of pointer to i: ", *ptr)
}
