package rectangle

import "math"

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
