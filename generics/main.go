package main

import "fmt"

func main() {
	var intslice = []int{1, 2, 3}
	fmt.Println(sumslice(intslice))

	var floatslice = []float32{1, 2, 3}
	fmt.Println(sumslice(floatslice))
}

func sumslice[T int | float32 | float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
