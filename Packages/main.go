package main

import (
	"fmt"
	"module/assistant"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main package")
	assistant.Towrite()

	erro := checkmail.ValidateFormat("luisabianca93@gmail.com")
	fmt.Println(erro)
}
