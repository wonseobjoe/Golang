/*
	blocks/main/guest.go
*/

package main

import (
	"blocks/archi/bottom"
	"fmt"
)

func PressBellByGuest(arg string) {
	//bottom.Greeting()
	bottom.GreetingWithHello(arg)
}

func main() {
	var your_name string
	fmt.Print("Say your name: ")
	fmt.Scanln(&your_name)

	PressBellByGuest(your_name)
}
