// Exploring for loop in go , which is one of the iterative statement available in go, but you can also use this like while style , let move ahead

package main

import "fmt"

func main(){

	// for loop but while style

	fmt.Println("For loop but while style")
	i:=1
	for i<=3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Classic for loop")
	// classical for loop
	for j:=1; j<=5; j++{
		fmt.Println(j)
	}

	// exploring range keyword in golang
	// Note -: Range always starts from 0th index

	fmt.Println("Exploring range keyword in go")
	for k:= range 6 {
		fmt.Println(k)
	}

	// Exploring break and continue keyword in golang

	fmt.Println("Break")

	for l:=range 12{
		if l==6 {
			break
		}
		fmt.Println(l)
	}

	// Exploring continue keyword in go
	fmt.Println("Continue")

	for m:=range 12{
		if m==8 {
			continue
		}
		fmt.Println(m)
	}
}