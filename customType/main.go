package customType

import "fmt"

type customString string
func (text customString) logCustomString(){
  fmt.Println(text)
}

func CustomTypeMainFn(){
  var name customString = "Abhishek"
  name.logCustomString()
}
