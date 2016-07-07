package main

import "fmt"

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	fmt.Printf("value is : %d \n", c())
}
