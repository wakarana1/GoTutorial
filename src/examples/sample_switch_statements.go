package main

import (
  "fmt"
)

func number() int {
  num := 15 * 5
  return num
}

func main() {

  // SIMPLE SWITCH STATEMENT WITH DEFAULT
  // finger := 4
  // switch finger {
  // case 1:
  //   fmt.Println("thumb")
  // case 2:
  //   fmt.Println("index")
  // case 3:
  //   fmt.Println("middle")
  // case 4:
  //   fmt.Println("ring")
  // case 5:
  //   fmt.Println("pinky")
  // default:
  //   fmt.Println("Incorrect finger number")
  // }

  // MULTIPLE CASE SWITCH STATEMENT
  // letter := "i"
  // switch letter {
  // case "a", "e", "i", "o", "u":
  //   fmt.Println("letter is a vowel")
  // default:
  //   fmt.Println("letter is a consonant")
  // }

  //EXPRESSIONLESS SWITCH
  // num := 75
  // switch {
  // case num >= 0 && num <= 50:
  //   fmt.Println("Number is between 0 and 50")
  // case num >=51 && num <= 100:
  //   fmt.Println("Number is between 51 and 100")
  // case num >= 101:
  //   fmt.Println("number is greater than 100")
  // }

  //FALLTHROUGH SWITCH CASE
  switch num := number(); {
  case num < 50:
    fmt.Printf("%d is less than 50\n ", num)
    fallthrough
  case num < 100:
    fmt.Printf("%d is less than 100\n ", num)
    fallthrough
  case num < 200:
    fmt.Printf("%d is less than 200\n ", num)
  }
}
