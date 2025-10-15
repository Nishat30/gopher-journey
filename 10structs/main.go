package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	//no inheritance in golang, No super or parent
	nishat:= User{"nishat", "abc.dev", true,16}
    fmt.Println(nishat)
	fmt.Printf("nishat details are : %+v\n",nishat)
	
	fmt.Printf("my name is: %v\n",nishat.Name)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}
