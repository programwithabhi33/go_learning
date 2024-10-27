package structType

import (
	"errors"
	"fmt"
)

type Programmer struct {
	FirstName                    string
	lastName                     string
	age                          string
	hobby                        string
	favouriteProgrammingLanguage string
}

func newProgrammer(userInputFirstName, userInputLastName, userInputAge, userInputHobby, userInputFavouriteProgrammingLanguage string) (*Programmer, error) {
	if userInputFirstName == "" || userInputLastName == "" || userInputAge == "" {
		return nil, errors.New("Firstname, Lastname or Age should not be empty")
	}
	return &Programmer{
		FirstName:                    userInputFirstName,
		lastName:                     userInputLastName,
		age:                          userInputAge,
		hobby:                        userInputHobby,
		favouriteProgrammingLanguage: userInputFavouriteProgrammingLanguage,
	}, nil
}

func StructMainFn() {
	userInputFirstName := outputTextAndReturnUserInput("Enter first name")
	userInputLastName := outputTextAndReturnUserInput("Enter last name")
	userInputAge := outputTextAndReturnUserInput("Enter your age")
	userInputHobby := outputTextAndReturnUserInput("Enter your hobby")
	userInputFavouriteProgrammingLanguage := outputTextAndReturnUserInput("Enter your favourite programming language")

	var programmerData *Programmer
	programmerData, err := newProgrammer(userInputFirstName, userInputLastName, userInputAge, userInputHobby, userInputFavouriteProgrammingLanguage)
	if err != nil {
		fmt.Println(err)
		return
	}

	programmerData.outputProgrammerDetails()
	programmerData.clearProgrammerName()
	programmerData.outputProgrammerDetails()
}

func outputTextAndReturnUserInput(outputText string) string {
	fmt.Println(outputText)
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}

func (user Programmer) outputProgrammerDetails() {
	//You access like this as well
	//user.firstName
	//But the technically correct to access the methods and properties on a struct pointer is
	//(*[STRUCT_POINTER]).[FIELD_NAME] in this case e.g,. (*user).firstName

	fmt.Println("programmer first name is:", user.FirstName)
	fmt.Println("programmer last name is:", user.lastName)
	fmt.Println("programmer last age is:", user.age)
	fmt.Println("programmer last hobby is:", user.hobby)
	fmt.Println("programmer favourite language is:", user.favouriteProgrammingLanguage)
}
func (user *Programmer) clearProgrammerName() {
	user.FirstName = ""
	user.lastName = ""
}
