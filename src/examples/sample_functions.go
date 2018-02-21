package main

import (
  "fmt"
)

// SIMPLE FUNCTION
// func calculateBill(price, no int) int {
//   var totalPrice  = price * no
//   return totalPrice
// }
// func main() {
//   price, no := 90, 6
//   totalPrice := calculateBill(price, no)
//   fmt.Println("Your total is", totalPrice)
// }

// MULTIPLE RETURN FUNCTION
func rectangleProps(length, width float64)(area, perimeter float64) {
  area = length * width
  perimeter = (length + width) * 2
  return // NO EXPLICIT RETURN VALUE BECAUSE IT'S ALREADY DECLARED
}
// func main() {
//   area, perimeter := rectangleProps(11.5, 12.6)
//   fmt.Printf("Area of rectangle %f, Perimeter of rectangle %f", area, perimeter)
// }
func main() {
  area, _ := rectangleProps(11.5, 12.6) // USE BLANK IDENTIFIER _ WHEN YOU DON'T NEED OTHER PARTS OF FUNCITON
  fmt.Printf("Area of rectangle %f", area)
}
