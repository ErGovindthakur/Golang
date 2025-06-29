package main

import "fmt"

func main(){
	fmt.Println("Welcome to our Go Todo Application !")

	var taskList = []string{"Code","Practice"}

	fmt.Println()
	fmt.Println("Previous Task -: ")
	showTask(taskList)

	fmt.Println()
	taskList = addTask(taskList,"Eat")
	taskList = addTask(taskList,"Sleep")


	fmt.Println()
	fmt.Println("Updated Task -: ")
	showTask(taskList)

}

// 1. show all task
func showTask(taskList[]string){
	for index, task:=range taskList{
		fmt.Printf("%d. %s\n",index+1,task)
	}
}

// 2. create all task
func addTask(taskList[]string,task string) []string{
	var newTask = append(taskList,task)
	return newTask
}