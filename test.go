package main

import "fmt"

func main() {
	// //Basic Syntax
    // fmt.Println("Hello world")// println for new line
	// fmt.Print("Hello")// print for same line

	// //Declaring variables
	// var num1=2
	// var num2=3
	// fmt.Print(num1+num2)

	// //shortway using :=
	// age:=21
	// fmt.Println(age)

	// //multiple variable assigning
	// x,y,z:=10,20,30
	// fmt.Println(x,y,z)

	// //constants
	// const age=21
	// fmt.Println(age)
	// // age=22 //throws error that cannot assign value to this

	// //Loops in GoLang->Only one loop that is For loop
	// //1.Infinite loop
	// for{
	// 	fmt.Println("Hello")
	// }

	// //2.Specifying condition
	// i:=1
	// for i<=5{
	// 	fmt.Println("Hello",i)
	// 	i++
	// }

	//3.condition in single line
	for i:=1;i<=3;i++{
		fmt.Println("Hello",i)
	}
}
