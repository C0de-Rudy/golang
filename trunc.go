package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter Floating Point Number: ")
	scanner.Scan()
	input, _ := strconv.ParseFloat(scanner.Text(), 64)
	input = math.Floor(input)
	fmt.Printf("Floating Point Number Truncated: %.f", input)

}
