package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Input String: ")
	scanner.Scan()
	input := (scanner.Text())
	input = strings.ToLower(input)

	if strings.HasPrefix(input, "i") &&
		strings.Contains(input, "a") &&
		strings.HasSuffix(input, "n") {

		fmt.Printf("Found!")
	} else {

		fmt.Printf("Not Found!")
	}

}
