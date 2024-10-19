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
  programmerData.outputProgrammerDetails()
  programmerData.clearProgrammerName()
  programmerData.outputProgrammerDetails()
}

func outputTextAndReturnUserInput(outputText string) string {
	fmt.Println(outputText)
	var userInput string 
	fmt.Scan(&userInput)
	return userInput
}

func (user programmer) outputProgrammerDetails(){
  //You access like this as well
  //user.firstName
  //But the technically correct to access the methods and properties on a struct pointer is
  //(*[STRUCT_POINTER]).[FIELD_NAME] in this case e.g,. (*user).firstName

  fmt.Println("programmer first name is:", user.firstName)
  fmt.Println("programmer last name is:", user.lastName)
  fmt.Println("programmer last age is:", user.age)
  fmt.Println("programmer last hobby is:", user.hobby)
  fmt.Println("programmer favourite language is:", user.favouriteProgrammingLanguage)
}
func (user *programmer) clearProgrammerName(){
  user.firstName = ""
  user.lastName = ""
}
