package main

import "fmt"

// public variable (first later should be capital)
const LoginTocken string = "kj$3r34ds99bcsDD4"

func main() {
	fmt.Printf("Variables: ")

	var username string = "name1"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)
	fmt.Printf("\n")

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)
	fmt.Printf("\n")

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T \n", smallVal)
	fmt.Printf("\n")

	var smallFloat float32 = 255.55656569989374
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)
	fmt.Printf("\n")

	var bigFloat float64 = 255.55656569989374
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type %T \n", bigFloat)
	fmt.Printf("\n")

	// by default variable has 0 as a value
	var justTheInt int
	fmt.Println(justTheInt)
	fmt.Printf("Variable is of type %T \n", justTheInt)
	fmt.Printf("\n")

	// implicit type
	var myName = "Ishan Jawade"
	fmt.Println(myName)
	fmt.Printf("Variable is of type %T \n", myName)
	fmt.Printf("\n")

	// no var keywaord style declaration ( := )
	// only allow inside the function
	totalBalance := 50300.85
	fmt.Println(totalBalance)
	fmt.Printf("Variable is of type %T \n", totalBalance)
	fmt.Printf("\n")

	fmt.Println(LoginTocken)
	fmt.Printf("Variable is of type %T \n", LoginTocken)
	fmt.Printf("\n")

}
