package main

import "fmt"

func main() {
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 12 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a kid")
	}
	f1()
	f2()
}

// we could also initilize a variable
// inside of if

func f1() {
	if age := 20; age >= 18 {
		fmt.Println("You are an adult !")
	}
	// This will give an error
	// fmt.Println("Your age is : ", age)
}

// for loops inside go
func f2() {
	for i := 0; i <= 5; i++ {
		fmt.Println("The numerber is : ", i)
	}
}

// infinite loops
func f3() {
	for {
		fmt.Println("This is infinite loop ....")
	}
}

// Looping over index/slices
func f4() {
	nums := []int{2, 4, 6, 7}
	for index, value := range nums {
		fmt.Println("Index : ", index, "Value : ", value)
	}
}

// Switch statement is powerful in go
// no need to break after each case

func f5() {
	day := 4

	// Switch with a variable
	switch day {
	case 1:
		fmt.Println("Day 1")
	case 2:
		fmt.Println("Day 2")
	case 4:
		fmt.Println("Day 3")
	}

	// switch without a variable
	// is like a if-else
	switch {
	case day == 1:
		fmt.Println("It is Monday")
	case day == 4:
		fmt.Println("It is Thursday")
	}
}

// Defer postpones the function execution
// function untill surrounding function exists
// Multiple defer calls execute in LIFO order

func f6() {
	defer fmt.Println("Third")
	defer fmt.Println("Second")
	fmt.Println("First")
}

// Functions with multiple return values
func swap(a int, b int) (int, int) {
	return b, a
}

// Veradic function can take any
// number of inputs

func f7(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func f7_help() {
	val := f7(1, 2, 3, 4, 5)
	fmt.Println("Sum is : ", val)

	// passing a slic
	arr := []int{1, 2, 3}
	nval := f7(arr...)
	fmt.Println("Sum is : ", nval)
}
