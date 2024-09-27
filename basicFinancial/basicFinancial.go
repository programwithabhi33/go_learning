package basicFinancial

import "fmt"

func ShowChoices(){
  accountBalance := 1000.0
  var choice int
  fmt.Println("Welcome to basic financial application!")
  fmt.Println("What do you want to do!")
  fmt.Println("1. Check balance")
  fmt.Println("2. Deposit")
  fmt.Println("3. Withdraw money")
  fmt.Println("4. Exit")
  
  fmt.Print("Your choice: ")
  fmt.Scan(&choice) 

  if choice == 1 {
    fmt.Println("Your account balance is:", accountBalance)
  }else if choice == 2 {
    var depositAmount float64
    fmt.Println("Your deposit:")
    fmt.Scan(&depositAmount)
    accountBalance += depositAmount
    fmt.Println("Balance updated!, New account balance is:", accountBalance)

  }else if choice == 3 {
    var withdrawalAmount float64
    fmt.Println("Withdrawal amount:")
    fmt.Scan(&withdrawalAmount)
    accountBalance -= withdrawalAmount 
    fmt.Println("Withdraw successfull!, New balance is:", accountBalance)
  }else{
    fmt.Println("You choose to exit from application!")
  }
}
