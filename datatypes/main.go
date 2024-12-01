package main

import "fmt"

func main() {
	var a int16 = 5
	var b int32 = 8
	c := a + int16(b)
	fmt.Println(c)

	str := "Yo!"
	fmt.Println(str)

	const pi = 3.14
	var isit bool = true
	fmt.Println(isit, pi)
}
