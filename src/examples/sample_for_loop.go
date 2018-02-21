package main

import (
  "fmt"
)

func main() {
  //SIMPLE FOR LOOP
  // for i := 1; i <= 10; i++ {
  //   fmt.Printf("%d ", i)
  // }

  // FOR LOOP WITH BREAK
  // for i := 1; i <= 10; i++ {
  //   if i > 5 {
  //     break //loop is terminated if i > 5
  //   }
  //   fmt.Printf("%d ", i)
  // }
  // fmt.Printf("\nline after for loop")

  // CONTINUE WILL SKIP THE REST OF THE FUNCTION
  // for i := 1; i <= 10; i++ {
  //   if i%2 == 0 {
  //     continue
  //   }
  //   fmt.Printf("%d ", i)
  // }

  // SEMICOLONS OMITTED AND ONLY THE CONDITION EXISTS
  // i := 0
  // for i <= 10 {
  //   fmt.Printf("%d ", i)
  //   i += 2
  // }

  //multiple initialisation and increment
  for no, i := 10, 1; i <= 10 && no <= 19; i, no = i + 1, no + 1 {
    fmt.Printf("%d * %d = %d\n", no, i, no * i)
  }

}
