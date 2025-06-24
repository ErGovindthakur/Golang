package main

import "fmt"

func data(num int)int{

	fmt.Println("Here is Reversed number code -: ")
	
	var rev int;
	for num!=0{
		var rem = num % 10
		rev = (rev*10)+rem
		num = num/10
	}
	return rev
}