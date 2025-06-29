package main

import "fmt"

func main(){
	// Data Types
	const myName string = "Govind" // 1. string 
	fmt.Println(myName)

	var age int = 20; // 2. integer (number)
	fmt.Println(age)


	// variable => used to store the types of data

	const myStatement = "Learn and Grow"; // here store string value into "myStatement"
	fmt.Println(myStatement)

	isTrue := true
	fmt.Println(isTrue)


	// In go collection of similar data types is store in 1. slice, 2. Array

	// 1. slice => it stores collection of similar types of data, but dynamic in size
	var taskItem = []string{"Study","play","eat","sleep","practice"};
	fmt.Println(taskItem)

	// 2. Array => it is same like slice but fixed in size

	// alternate to define variables
	 weekEnd := [2]string{"Saturday","Sunday"}
	 fmt.Println(weekEnd)

}