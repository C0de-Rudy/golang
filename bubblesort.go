package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 10)
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	defer fmt.Println("Integer Limit Reached\n PROGRAM TERMINATED")

	for i := 0; i < 10; i++ {
		fmt.Printf("Enter Number or type 'Exit' to Terminate: ")
		if scanner.Scan() {
			input = scanner.Text()
		}
		if input == "Exit" {
			fmt.Println("PROGRAM TERMINATED")
			return
		}
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Input: ", err)
			continue
		}
		slice = append(slice, value)
		sort.Ints(slice)

		fmt.Println(BubbleSort(slice))
	}
}

func BubbleSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
