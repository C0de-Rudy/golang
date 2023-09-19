package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func getArray() []int {
	stringInput := scanInput()
	stringSlice := strings.Fields(stringInput)
	return arrConvToInt(stringSlice)
}

func scanInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Please input up to 20 integers, seperated with spaces.")
		if scanner.Scan() {
			return scanner.Text()
		} else {
			fmt.Printf("Input error: %s, please try again.", scanner.Err().Error())
			continue
		}
	}
}

func arrConvToInt(array []string) []int {
	intArr := make([]int, 0, 20)
	for _, value := range array {
		inputInt, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		intArr = append(intArr, inputInt)
	}
	return intArr
}

func BubbleSort(slice []int, c chan []int, wg *sync.WaitGroup) {
	fmt.Println(slice)
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	wg.Done()
	c <- slice
}

func MergeSort(a []int, b []int, c chan []int, wg *sync.WaitGroup) {

	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	wg.Done()
	c <- final
}

func main() {
	array := getArray()
	var slices [][]int
	chanMain := make(chan []int, 4)
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {

		min := (i * len(array) / 4)
		max := ((i + 1) * len(array)) / 4

		slices = append(slices, array[min:max])
	}
	wg.Add(4)
	go BubbleSort(slices[0], chanMain, &wg)
	go BubbleSort(slices[1], chanMain, &wg)
	go BubbleSort(slices[2], chanMain, &wg)
	go BubbleSort(slices[3], chanMain, &wg)
	wg.Wait()

	a := <-chanMain
	b := <-chanMain
	c := <-chanMain
	d := <-chanMain

	wg.Add(3)
	go MergeSort(a, b, chanMain, &wg)
	go MergeSort(c, d, chanMain, &wg)

	e := <-chanMain
	f := <-chanMain

	go MergeSort(e, f, chanMain, &wg)
	wg.Wait()
	fmt.Println(<-chanMain)
}
