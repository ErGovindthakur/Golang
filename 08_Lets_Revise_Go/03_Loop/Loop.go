package main

import (
	"fmt"
)

func main() {
	// creating a slice and array and iterating with loop

	var carList = []string{"Audi", "Mercedes", "Range-Rover"}

	for index, list := range carList {
		// fmt.Println(index+1,".",list)

		// formatted type
		fmt.Printf("%d. %s\n", index+1, list)
	}
}
