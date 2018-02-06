package main

import "fmt"

func main() {
  // var age int = 28

  // var age = 28 inferred variable

  // var width, height int = 100, 50 declaring multiple variables

  var width, height = 100, 50 // declaring multiple inferred variables

  // var (
  //   name = "Chris"
  //   age = 28
  //   height int
  //   ) declaring different variables with different inferred value types

  name, age := "chris", 28 //shorthand declaration for new variables only


  fmt.Println("My name is ", name, "My age is", age)
  fmt.Println("This box's dimensions are" , width, height)
}
