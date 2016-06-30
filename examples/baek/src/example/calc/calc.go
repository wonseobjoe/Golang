package calc

import(
	"fmt"
)


func Plus(x int , y int) int {
	return x+y
}

func Minus(x int , y int) int {
	return x-y
}

func Multiply(x int , y int) int {
	return x*y
}

func Divide(x float64,y float64  ) float64 {
	return x/y
}

func Execute(){
	
	fmt.Println(" result plus : ",Plus(11, 12))
	fmt.Println(" result Minus : ",Minus(31, 12))
	fmt.Println(" result Multiply : ",Multiply(5, 3))
	fmt.Println(" result Divide : ",Divide(100, 3))
	pluss()
	
}

func pluss(){
	fmt.Println(" pluss private function")
}



