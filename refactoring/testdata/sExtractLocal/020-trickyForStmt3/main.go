package main

import (
	"fmt"
	"math"
)

func main() {
	x := 5.0
	for x < 10 { // <<<<< extractLocal,10,5,10,11,newVar,fail
		x++
	}
	for x > 10 {
		x++
	}
	for x+3 < 10 {
		x++
	}
	if math.Mod(x, 5) == 0 {
		fmt.Println("divisible by 5:")
	}
}
