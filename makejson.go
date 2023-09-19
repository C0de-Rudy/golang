package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\nInput Name Here: ")
	scanner.Scan()
	name := (scanner.Text())

	fmt.Printf("\nInput Address Here: ")
	scanner.Scan()
	address := (scanner.Text())

	{
		Mapmain := map[string]string{name: address}

		jsonmain, err := json.Marshal(Mapmain)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			fmt.Println("\nJSON Object: ", string(jsonmain))
		}

	}
}
