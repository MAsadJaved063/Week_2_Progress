package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var TotalBalance float64
	nickles := 0.05
	dimes := 0.10
	quarters := 0.25
	for TotalBalance < 20 {
		switch rand.Intn(3) {
		case 0:
			TotalBalance += nickles
		case 1:
			TotalBalance += dimes
		case 2:
			TotalBalance += quarters
		}
		fmt.Printf("$%5.2f\n", TotalBalance)
	}

}
