package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MaxpriceRice = 6
const MaxpricePaneer = 4

func main() {
	var riceChannel = make(chan string)
	var paneerChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}

	for i := range websites {
		go checkRicePrices(websites[i], riceChannel)
		go checkPaneerPrices(websites[i], paneerChannel)
	}

	sendMessage(riceChannel, paneerChannel)
}

func checkRicePrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var price = rand.Float32() * 20
		if price < MaxpriceRice {
			c <- website
			break
		}
	}
}

func checkPaneerPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var price = rand.Float32() * 20
		if price < MaxpricePaneer {
			c <- website
			break
		}
	}
}

func sendMessage(r chan string, p chan string) {
	select {
	case website := <-r:
		fmt.Println("Found a deal on rice at: ", website)
	case website := <-p:
		fmt.Println("Found a deal on panner at: ", website)
	}
}
