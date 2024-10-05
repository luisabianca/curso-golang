package main

import "fmt"

func main() {
	var variable1 = "Variable 1" // Explicit Variable
	variable2 := "Variable 2"    // Implicit Variable with type inference

	var (
		person1 string = "Bianca"
		person2 string = "Raquel"
	)

	thing1, thing2 := "cellphone", "laptop"

	fmt.Println(variable1)
	fmt.Println(variable2)
	fmt.Println(person1, person2)
	fmt.Println(thing1, thing2)
}
