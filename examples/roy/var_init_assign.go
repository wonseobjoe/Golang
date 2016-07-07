package main

import (
	"fmt"
	_ "time"
)

func main() {
	var x, y int = 1, 2
	var age, name = 39, "Roy Joe"
	a, b, c, d := 1, 2, "String", false

	var (
		k, j         int = 10, 20
		address, sex     = "aaa", "male"
		unused       int = -1
		_                = unused
	)
	fmt.Println(x, y, age, name, a, b, c, d, k, j, address, sex, unused)
}
