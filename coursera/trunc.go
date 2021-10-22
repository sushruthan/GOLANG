package main

import (
	"fmt"
	"log"
)

var inputNumber float64

func main() {
	fmt.Println("Enter a Float number:=>")
	_, err := fmt.Scan(&inputNumber)
	if err != nil {
		log.Printf("[Error] Invalid input !")
	} else {
		fmt.Printf("The number you've entered is '%v'.\n", inputNumber)
		fmt.Printf("Truncated version of '%v' is '%v'.\n", inputNumber, int64(inputNumber))
	}

}
