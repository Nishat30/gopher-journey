package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter the rating for our pizza:")

	//comma of || err ok 

	input, _ := reader.ReadString('\n') //first part try second part catch
	fmt.Println("thank for rating, ",input)
	fmt.Printf("type of the rating is %T ",input)
}
