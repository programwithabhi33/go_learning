package pointer

import "fmt"

func PointerMainFunction(){
  //fmt.Println("Welcome to pointer's in GO")
  age := 45

  var agePointer *int
  //var anotherPointer *int
  agePointer = &age

  fmt.Println("Value stored in age pointer is:", *agePointer)
  //fmt.Println("Another pointer is:", anotherPointer)
  adultYears := getAdultYears(agePointer)
  fmt.Println("Adult years is", adultYears)
}

func getAdultYears(pointer *int) int {
  return *pointer - 18
}
