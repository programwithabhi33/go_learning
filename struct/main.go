package structType

import (
	"errors"
	"fmt"
)

type Programmer struct {
	firstName                    string
	lastName                     string
	age                          string
	hobby                        string
	favouriteProgrammingLanguage string
}
type Admin struct {
  username string
  password string
  Programmer
}

func NewAdmin(username, password string) Admin{
  return Admin{
    username: username,
    password: password,
    Programmer: Programmer{
      firstName: "ADMIN",
      lastName: "ADMIN",
      age: "---",
    },
  }
}

func New(userInputFirstName, userInputLastName, userInputAge, userInputHobby, userInputFavouriteProgrammingLanguage string) (*Programmer, error) {
	if userInputFirstName == "" || userInputLastName == "" || userInputAge == "" {
		return nil, errors.New("Firstname, Lastname or Age should not be empty")
	}
	return &Programmer{
		firstName:                    userInputFirstName,
		lastName:                     userInputLastName,
		age:                          userInputAge,
		hobby:                        userInputHobby,
		favouriteProgrammingLanguage: userInputFavouriteProgrammingLanguage,
	}, nil
}

func OutputTextAndReturnUserInput(outputText string) string {
	fmt.Println(outputText)
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}

func (user Programmer) OutputProgrammerDetails() {
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
func (user *Programmer) ClearProgrammerName() {
	user.firstName = ""
	user.lastName = ""
}
