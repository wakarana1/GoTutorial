package main

import (
  "fmt"
)

// FUNC
// func changeLocal(num [5]int) {
//   num[0] = 55
//   fmt.Println("inside function ", num)
// }

// MULTIDIMENSIONAL ARRAYS
// func printarray(a [3][2]string) {
//   for _, v1 := range a {
//     for _, v2 := range v1 {
//       fmt.Printf("%s ", v2)
//     }
//     fmt.Printf("\n")
//   }
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

  // RANGE RETURNS BOTH THE INDEX AND VALUE
  // a := [...]float64{67.5, 89.8, 21, 24}
  // sum := float64(0)
  // //for _, v := range a { // IGNORES THE INDEX
  // for i, v := range a {
  //   fmt.Printf("%d the element of a is %.2f\n", i, v)
  //   sum += v
  // }
  // fmt.Println("\nsum of all elements of a", sum)

  // MULTIDIMENSIONAL ARRAYS
  // a := [3][2]string {
  //   {"lion", "tiger"},
  //   {"cat", "dog"},
  //   {"pigeon", "peacock"},
  // }
  // printarray(a)
  // var b [3][2]string
  // b[0][0] = "apple"
  // b[0][1] = "samsung"
  // b[1][0] = "microsoft"
  // b[1][1] = "google"
  // b[2][0] = "at&t"
  // b[2][1] = "t-mobile"
  // fmt.Printf("\n")
  // printarray(b)

  // SLICE ARRAY
  // a := [5]int{76, 77, 78, 79, 80}
  // var b []int = a[1:4] //creates a slice from a[1] to a[3]
  // fmt.Println(b)
  // c := []int{6, 7, 8} //creates an array and returns a slice reference
  // fmt.Println(c)

  // MODIFY A SLICE
  // darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
  // dslice := darr[2:5]
  // fmt.Println("array before", darr)
  // for i := range dslice {
  //   dslice[i]++ //increase int by 1
  // }
  // fmt.Println("array after", darr)

  // MODIFY A SLICE 2
  // numa := [3]int{78, 79, 80}
  // nums1 := numa[:]
  // nums2 := numa[:]
  // fmt.Println("array before change 1", numa)
  // nums1[0] = 100
  // fmt.Println("array after modification to slice nums1", numa) // modifies the original slice of the array
  // nums2[1] = 101
  // fmt.Println("array after modification to slice nums2", numa) // modifies the original slice of the array

  // LENGTH AND CAPACITY OF ARRAY
  // fruitarray := [...]string{"apple", "orange", "grape", "mango", "watermelon", "pineapple", "chikoo"}
  // fruitslice := fruitarray[1:3]
  // fmt.Println("Array", fruitarray)
  // fmt.Println("Slice",fruitslice)
  // fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))
  // fruitslice = fruitslice[:cap(fruitslice)] //reslices fruitslice until its capacity
  // fmt.Println("After re-slicing length is", len(fruitslice))
  // fmt.Println("New slice", fruitslice)

  // CREATING SLICE USING MAKE
  // i := make([]int, 5, 5) // func make([]T, len, cap) where capacity is optional and will default to len
  // fmt.Println(i)

  // APPENDING TO A SLICE
  cars := []string{"Ferrari", "Honda", "Ford"}
  fmt.Println("cars: ", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
  cars = append(cars, "Toyota") // func append(s []T, x ...T) []T where x ...T acceps variable number of arguments for the param x (variadic function)
  fmt.Println("cars: ", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
}
