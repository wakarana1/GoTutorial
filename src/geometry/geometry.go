package main

import (
  "fmt"
  "geometry/rectangle" // IMPORTING CUSTOM PACKAGE
  "log"
)

func main() {
  var rectLen, rectWidth float64 = 6, 7
  fmt.Println("GEometrical shape properties")
    /* Area function of rectangle package used
    */
  fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
    /* Diagonal function of rectangle package
    */
  fmt.Printf("diagonal of the rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
}
