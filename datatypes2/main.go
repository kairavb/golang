package main

import "fmt"

func main() {
	var arr = [3]int{3, 4, 5}

	fmt.Println(arr[0])
	fmt.Println(arr[1])

	var arr1 [2]int
	fmt.Println(arr1)

	var amap = make(map[string]int)
	amap["cat"] = 1
	amap["dog"] = 2

	for name, num := range amap {
		fmt.Printf("Name: %v Num: %v\n", name, num)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("*")
	}
	fmt.Println()

	for i, v := range arr1 {
		fmt.Printf("index: %v\nvalue: %v\n", i, v)
	}

	// while loop
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i += 2
	}
}
