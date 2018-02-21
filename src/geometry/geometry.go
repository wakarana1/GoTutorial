package main
/* import throws and error because it's not being used so using instance _ silences it
  import (
    "geometry/rectangle"
  )
var _ = rectangle.Area // error silencer
*/

/* if you need to import a package but don't need to use any functions
  import (
    _ "geometry/rectangle"
  )
*/

import (
  "fmt"
  "geometry/rectangle" // IMPORTING CUSTOM PACKAGE
  "log"
)
/*
 * 1. package variables
*/
 var rectLen, rectWidth float64 = 6, 7
 /*
 * 2. init function to check if length and width are greater than zero
 */

 func init() {
   Println("main package initialized")
   if rectLen < 0 {
    log.Fatal("length is less than zero") // terminates execution if true
   }
   if rectWidth < 0 {
    log.Fatal("width is less than zero") // terminates execution if true
   }
 }

func main() {
  fmt.Println("Geometrical shape properties")
    /* Area function of rectangle package used
    */
  fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
    /* Diagonal function of rectangle package
    */
  fmt.Printf("diagonal of the rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
}
