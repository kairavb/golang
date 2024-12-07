package main

import "fmt"

func main() {
	var target int
	var arr = [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(arr)
	fmt.Println("Enter Target")
	fmt.Scan(&target)
	fmt.Println(binarySearch(&arr, target))
}

func binarySearch(arr *[15]int, t int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		pivot := (low + high) / 2
		if arr[pivot] > t {
			high = pivot - 1
		} else if arr[pivot] < t {
			low = pivot + 1
		} else {
			return pivot
		}
	}
	return -1
}
