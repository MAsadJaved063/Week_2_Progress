package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nickles := 5
	dimes := 10
	quarter := 25
	var TotalBalance = 0

	for TotalBalance < 2000 {
		switch rand.Intn(3) {
		case 0:
			TotalBalance += nickles
		case 1:
			TotalBalance += dimes
		case 2:
			TotalBalance += quarter
		}

		BalanceInDollars := TotalBalance / 100
		cents := TotalBalance % 100
		fmt.Printf("$%v.%02v\n", BalanceInDollars, cents)
	}

}
