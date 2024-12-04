package main

import "fmt"

func main() {
	var intslice = []int{1, 2, 3}
	fmt.Println(sumslice(intslice))

}

func sumslice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
