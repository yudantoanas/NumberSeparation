package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// test number
	number := 1345679

	// convert number to string
	numString := strconv.Itoa(number)

	for i := 0; i < len(numString); i++ {
		// print
		fmt.Println(string(numString[i]) + strings.Repeat("0", len(numString)-(i+1)))
	}
}
