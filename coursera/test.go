package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var enteredString, loweredString string

	fmt.Print("Enter string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	enteredString = scanner.Text()
	loweredString = strings.ToLower(enteredString)

	if strings.HasPrefix(loweredString, "i") && strings.ContainsRune(loweredString, 'a') && strings.HasSuffix(loweredString, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
