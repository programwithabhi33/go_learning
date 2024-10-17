package structType

import "fmt"

type programmer struct {
	firstName                    string
	lastName                     string
	age                          string
	hobby                        string
	favouriteProgrammingLanguage string
}

func StructMainFn() {
	userInputFirstName := outputTextAndReturnUserInput("Enter first name")
	userInputLastName := outputTextAndReturnUserInput("Enter last name")
	userInputAge := outputTextAndReturnUserInput("Enter your age")
	userInputHobby := outputTextAndReturnUserInput("Enter your hobby")
	userInputFavouriteProgrammingLanguage := outputTextAndReturnUserInput("Enter your favourite programming language")

	var programmerData programmer
	programmerData = programmer{
		firstName:                    userInputFirstName,
		lastName:                     userInputLastName,
		age:                          userInputAge,
		hobby:                        userInputHobby,
		favouriteProgrammingLanguage: userInputFavouriteProgrammingLanguage,
	}
  fmt.Println("Programmer Data is:", programmerData)
}

func outputTextAndReturnUserInput(outputText string) string {
	fmt.Println(outputText)
	var userInput string 
	fmt.Scan(&userInput)
	return userInput
}
