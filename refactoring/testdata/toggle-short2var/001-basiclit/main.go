// <<<<< toggle,9,3,9,8,pass
package main

import "fmt"
import "reflect"

// Test for changing i := 3 to a var declaration
func main() {
  i := 3
  fmt.Println("type of i is", reflect.TypeOf(i))
  fmt.Println("value of i is", i)
}
