package main

import (
	"fmt"
	"sort"
)

func main() {
	temp := 1
	var sli = make([]int, 3)
	for temp > 0 {
		_, err := fmt.Scan(&temp)
		if err == nil {
			sli = append(sli, temp)
			sort.Ints(sli)
			fmt.Println(sli)
		} else {
			break
		}
	}
}
