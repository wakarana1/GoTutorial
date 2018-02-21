package main

import (
  "fmt"
)

func main() {
  // REGULAR IF ELSE STATEMENT
  num := 99
  if num < 50 { // checks if number is even
    fmt.Println("Number is less than or equal to 50")
  } else if num >= 51 && num <= 100 {
    fmt.Println("Number is between 51 and 100")
  } else {
    fmt.Println("Number is greater than 100")
  }

  // FUNCTION WITH STATEMENT INSIDE IF
  // if num := 10; num % 2 == 0 {
  //   fmt.Println("Even Number")
  // } else {
  //   fmt.Println("Odd Number")
  // }
}
