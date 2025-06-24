package main

import "fmt"

func main(){
	fmt.Println("Exploring conditional statement in go")

	var age int 

	fmt.Print("Enter the age -:")
	fmt.Scan(&age)

	if age > 18 && age <= 45{
		fmt.Println("You are adult")
	}else if age >= 13 && age <= 18{
		fmt.Println("You are teenager")
	}else{
		fmt.Println("You are a kid")
	}
}