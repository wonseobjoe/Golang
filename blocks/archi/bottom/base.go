/*
	blocks/archi/bottom/base.go
*/
package bottom

import (
	"fmt"
)

func Greeting() {
	fmt.Println("Hello, Anybody !!!")
}

func GreetingWithHello(arg string) {
	fmt.Println("Hello, " + arg + " !!!")
}
