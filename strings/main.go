package main

import "fmt"

func main() {
	str := "resume"
	fmt.Println(str)
	fmt.Println(str[0:5])
	fmt.Println(str[0])

	fmt.Printf("value %v, type %T\n", str[0], str[0])

	for i, v := range str {
		fmt.Println(i, v)
	}
	fmt.Println(len(str))

	var str1 string = "résumé"
	for i, v := range str1 {
		fmt.Println(i, v)
	}
	fmt.Println(len(str1))

	// []rune == []int32
	var str2 = []rune("résumé")
	for i, v := range str2 {
		fmt.Println(i, v)
	}
	fmt.Println(len(str))
}
