package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 2
	if a+b < c {
		fmt.Println("")
		a = 5
		b = 6
		if a+b < c { // <<<<< var,13,5,13,8,newVar,pass
			fmt.Println("this is the area of a rectangle: ")
			fmt.Println(a + b)
		}
	}
	d := a + b
	fmt.Println(d)
}
