package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int8 = -127         // INT positive and negative numbers.
	var negativeNumber uint8 = 127 // UINT only positive numbers.
	var number2 int32 = 33000

	//alias
	//INT32 = RUNE
	//UINT8 = BYTE
	var number3 rune = 300000000
	var number4 byte = 127

	fmt.Println(number)
	fmt.Println(negativeNumber)
	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Println(number4)

	const name string = "Bianca"
	const floatNumber float32 = 23.005
	var floatNumber2 float64 = 2323.00500000

	fmt.Println(name)
	fmt.Println(floatNumber)
	fmt.Println(floatNumber2)

	var bolean1 bool = true
	var bolean2 bool = false

	fmt.Println(bolean1, bolean2)

	//type erro
	var erro error
	fmt.Println(erro)

	//package erro
	var erro2 error = errors.New("Sorry, Something go wrong")
	fmt.Println(erro2)

}
