package rectangle

import "math"
import "fmt"

/* used to perform initialization tasks and can also be used to verify
   the correctness of the program before the execution starts
*/

func init() {
  fmt.Println("rectangle package initialized")
}

/* functions are capitalized because they're exported functions
   if they're not to capitalized, it will throw an error
*/
func Area(len, wid float64) float64 {
  area := len * wid
  return area
}

func Diagonal(len, wid float64) float64 {
  diagonal := math.Sqrt((len * len) + (wid * wid))
  return diagonal
}
