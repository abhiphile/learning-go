package main

import (
	"bufio"
	"fmt"
	"os"
)

func printing() {
	fmt.Println("Hello World") // Prints with a new line
	fmt.Print("Hello World\n") // Prints without a new line

	name := "Allice"
	age := 20

	// Prints formatted string like C language
	fmt.Printf("%s age is %d\n", name, age)

	/*
		Formatted String overview
		-------------------------
		%s	String
		%d	Integer
		%f	Float
		%.2f	Floating Point 2 decimal
		%t	Boolean eg-op true
		%v Default format
		%T Type of variable
	*/

	fmt.Println("Printing Type of variable")
	fmt.Printf("Type of name is : %T\n", name)
}

func scanning() {
	var name string
	var age int
	fmt.Println("Enter your name and age : ")
	/*
		fmt.Scan() Only works when input is given
		in a single line seperated by comma
	*/
	fmt.Scan(&name, &age) // Reads input

	/*
		fmt.Scanln() Scans for complete
		line but stops at first space
	*/
	fmt.Scanln(&name) // Reads untill space

	// Reads input like C structured input
	fmt.Sscanf("%s %d", name, age)

	/*
		Taking complete line input
		-------------------------
	*/
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered, ", input)
}

func main() {

}
