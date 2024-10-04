package basicFinancial

import "fmt"
import "os"

func storeBalanceInFile(balance float64){
  balanceText := fmt.Sprint(balance)
  os.WriteFile("./balance.txt", []byte(balanceText), 0644)
}

func ShowChoices() {
	accountBalance := 1000.0
	fmt.Println("Welcome to basic financial application!")
	for {
		var choice int
		fmt.Println("What do you want to do!")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		/*if choice == 1 {
			fmt.Println("Your account balance is:", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Println("Your deposit:")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount, the deposit amount should be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated!, New account balance is:", accountBalance)

		} else if choice == 3 {
			var withdrawalAmount float64
			fmt.Println("Withdrawal amount:")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount, the withdrawal amount should be greater than 0")
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount, you can't withdraw amount greater than you have in your account!")
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Withdraw successfull!, New balance is:", accountBalance)
		} else {
			fmt.Println("You choose to exit from application!")
			break
		}*/

		switch choice {
		case 1:
			fmt.Println("Your account balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Println("Your deposit:")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount, the deposit amount should be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated!, New account balance is:", accountBalance)
      storeBalanceInFile(accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Println("Withdrawal amount:")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount, the withdrawal amount should be greater than 0")
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount, you can't withdraw amount greater than you have in your account!")
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Withdraw successfull!, New balance is:", accountBalance)
      storeBalanceInFile(accountBalance)
		default:
			fmt.Println("You choose to exit from application!")
			fmt.Println("Thanks for using our financial application")
			return
		}
	}
}
