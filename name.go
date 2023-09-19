package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\nFile Name: ")
	scanner.Scan()
	filename := (scanner.Text())

	namesfile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		scanFile := bufio.NewScanner(namesfile)
		for scanFile.Scan() {
			line := scanFile.Text()
			newObject := strings.Fields(line)
			fmt.Println(newObject)
		}
	}

}
func intoStruct([]string) {

}
