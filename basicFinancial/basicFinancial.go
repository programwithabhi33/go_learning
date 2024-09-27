package basicFinancial

import "fmt"

func ShowChoices(){
  var choice int
  fmt.Println("Welcome to basic financial application!")
  fmt.Println("What do you want to do!")
  fmt.Println("1. Check balance")
  fmt.Println("2. Withdraw money")
  fmt.Println("3. Deposit")
  fmt.Println("4. Exit")
  
  fmt.Print("Your choice: ")
  fmt.Scan(&choice) 
  fmt.Println("Your choice is: ", choice)
}
