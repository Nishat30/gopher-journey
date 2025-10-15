package main

import "fmt"

//jwtToken:= 300000 //this will show erro coz we are not alllowed to use walrus opeartor globally or more precisely outside a method

const LoginToken string = "mytoken" //L capital so public

func main() {
	var username string = "nishat" //if we declare and dont use that var it will show error
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username) //%T for type

	var isLoggedIn bool = true //boolean
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 //int(numeric int)
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.2132434421312312 //int(numeric int)
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anothervar int
	fmt.Println(anothervar)
	fmt.Printf("variable is of type:%T\n", anothervar)

	//implicit type
	var website = "google.com"
	fmt.Println(website)

	//no var style
	numberOfUser := 300000 //:= (walrus operator)
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable type of logintoken is: %T\n", LoginToken)
}
