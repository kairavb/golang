package main

import "fmt"

type engine struct {
	mpg     uint8
	gallons uint8
	owner
}

type owner struct {
	name string
}

func (t engine) milesleft() uint8 {
	return t.gallons * t.mpg
}

func main() {
	var myEngine engine = engine{}
	myEngine.owner = owner{"kairav"}
	myEngine.mpg = 3
	myEngine.gallons = 10

	fmt.Println(myEngine)
	fmt.Printf("total miles left: %v\n", myEngine.milesleft())

}
