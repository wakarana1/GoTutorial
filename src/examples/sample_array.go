package main

import (
  "fmt"
)

// FUNC
// func changeLocal(num [5]int) {
//   num[0] = 55
//   fmt.Println("inside function ", num)
// }


func main() {
  // SHORTHAND DECLARATION OF AN INTEGER ARRAY WITH A LENGTH OF 3
  // a := [3]int{12 ,78 ,50}

  // ... MAKES THE COMPILER DETERMINE THE LENGTH OF THE ARRAY
  // a := [...]int{12, 78, 50}
  // fmt.Println(a)

  // PASSING THROUGH FUNCTION
  // num := [...]int{5, 6, 7, 8, 8}
  // fmt.Println("before passing to function ", num)
  // changeLocal(num)
  // fmt.Println("after passing through function ", num)

  // FINDING LEN OF ARRAY
  // a := [...]float64{67.5, 89.8, 21, 24}
  // fmt.Println("length of a is",len(a))

  // ITERATING THROUGH AN ARRAY WITH FOR LOOP
  // a := [...]float64{67.5, 89.8, 21, 24}
  // for i := 0; i < len(a); i++ {
  //   fmt.Printf("%d the element of a is %.2f\n", i, a[i])
  // }

  // range returns both the index and value
  a := [...]float64{67.5, 89.8, 21, 24}
  sum := float64(0)
  for i, v := range a {
    fmt.Printf("%d the element of a is %.2f\n", i, v)
    sum += v
  }
  fmt.Println("\nsum of all elements of a", sum)
}
