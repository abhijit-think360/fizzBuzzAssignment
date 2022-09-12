package main

import (
	"fmt"
)

func main() {

	fmt.Print("Enter lower limit: ")
	var lowerLimit int
	fmt.Scanln(&lowerLimit)

	fmt.Print("Enter upper limit: ")
	var upperLimit int
	fmt.Scanln(&upperLimit)

	for i := lowerLimit; i <= upperLimit; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(i int) {
	fizz := "fizz"
	buzz := "buzz"

	if i%3 == 0 && i%5 == 0 {
		fmt.Println(i, fizz+buzz)
	} else if i%3 == 0 {
		fmt.Println(i, fizz)
	} else if i%5 == 0 {
		fmt.Println(i, buzz)
	} else {
		fmt.Println(i)
	}
}
