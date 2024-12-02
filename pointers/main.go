package main

import "fmt"

func main() {
	// new gives random free memory and point the pointer to it
	var p *int = new(int)

	// here is a dangling pointer, which points to nothing or nil or null in memory
	// do not dereference a null pointer, always init a pointer by pointing at something
	// var p *int

	var i int
	fmt.Println(*p, i)

	// now pointer points to address of i
	p = &i

	// dereferencing the pointer
	*p = 1

	fmt.Println(*p, i)

	var arr = [5]int{1, 2, 3, 4, 5}

	result := square(&arr) // passed orignal, not as copy

	fmt.Println(result, arr)
}

func square(arr1 *[5]int) [5]int { // taken as pointer to orignal address
	for i := range arr1 {
		arr1[i] *= arr1[i]
	}
	return *arr1 // return dereferenced value
}
