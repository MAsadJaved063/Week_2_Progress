package main

import (
	"fmt"
	"math"
)

func main() {

	age := 41
	marsAge := float64(age)
	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("I am", marsAge, "years old on Mars.")

	//Convert types with caution
	var bh float64 = 32767
	var h = int16(bh)
	fmt.Println(h)

	v := 42
	if v >= 0 && v <= math.MaxUint8 {
		v8 := uint8(v)
		fmt.Println("converted:", v8) //Prints converted: 42

	}

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Print(string(pi), string(alpha), string(omega), string(bang))

	countdown := 9
	str := fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	fmt.Println(str)

	//Converting a Boolean to a string
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("\nReady for launch:", launchText)
	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)

}
