package pointer

import "fmt"

func PointerMainFunction(){
  //fmt.Println("Welcome to pointer's in GO")
  age := 45

  var agePointer *int
  agePointer = &age

  fmt.Print("age pointer is:", agePointer)
}
