package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	fmt.Println("enter a string")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	str = scanner.Text()
	fmt.Println("captured:", str)
	length := len(str) - 1
	if ((string(str[0]) == "i") || (string(str[0]) == "I")) && ((string(str[length]) == "n") || (string(str[length]) == "N")) {
		for i, j := range str {
			if i > 0 && ((string(j) == "A") || (string(j) == "a")) {
				fmt.Println("Found!")
				return
			}
		}
		println("Not Found!")
	} else {
		println("Not Found!")
	}
}
