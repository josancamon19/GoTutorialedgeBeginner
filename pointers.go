package main

import "fmt"

// In Go, when we call a function and pass in a bunch of arguments to that function,
// the language creates copies of the arguments which are then used within said function.
func main6() {
	a := 2
	addOne(&a)
	fmt.Println(a)

	// GO BY EXAMPLE
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// use the * symbol at the point at which we are declaring the variable and it will turn the variable into a pointer variable.
func addOne(a *int) {
	*a += 3
}
