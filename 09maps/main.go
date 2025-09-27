package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	fmt.Println("list of all languages: ", languages)
	fmt.Println("js shorts for: ", languages["js"])

	delete(languages, "rb")
	fmt.Println("list of all languages: ", languages)

	//loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("for key %v,value is %v\n", key, value)
	}
}
