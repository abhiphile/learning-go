package main

import "fmt"

/*
In Go variables are classified as their decleration method,
storage class, and scope
*/

// Decleration Methods

func f1() {
	// Explicitly decleared variables
	var name string = "Allice"
	var age int = 20
	fmt.Print(name, age)
}

func f2() {
	// Short declearation using :=
	name := "Bob"
	age := 20
	fmt.Print(name, age)
}

func f3() {
	// Multiple variable decleartion
	var a, b, c int = 3, 4, 5
	x, y := 8, "Hello"
	fmt.Print(a, b, c, x, y)
}

func f4() {
	// Block decleartion
	var (
		a    int    = 3
		name string = "Hello"
	)
	println(a, name)
}

// Scope

func f5() {
	// Local variable
	x := 20
	println(x)
}

// Global Variable
var weekNo = "Week 1"

const PI = 3.1415

/*
Type	Description						Example
----	------------					-------
int		Integer 						(platform-dependent) var x int = 10
int8	8-bit integer 					(-128 to 127) var x int8 = 127
int16	16-bit integer					var x int16 = 32767
int32	32-bit integer 					(rune)	var x int32 = 2147483647
int64	64-bit integer					var x int64 = 9223372036854775807
uint	Unsigned int 					(platform-dependent)	var x uint = 10
float32	32-bit floating point			var f float32 = 3.14
float64	64-bit floating point			var f float64 = 3.1415926535
bool	Boolean (true/false)			var flag bool = true
string	Sequence of characters			var name string = "Alice"



B. Derived Data Types
Type		Description					Example
----		-----------					-------
Pointer		Stores memory address		var p *int
Array		Fixed-size collection		var arr [5]int
Slice		Dynamic-size array			var sl []int
Map			Key-value pairs				var m map[string]int
Struct		Custom data type			type Person struct { Name string }
Interface	Abstract type				var i interface{}
Function	Function as a variable		var f func(int) int
*/

func f() {
	fmt.Println("Main")
	f1()
	f2()
	f3()
	f4()
	f5()
}
