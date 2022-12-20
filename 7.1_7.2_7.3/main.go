package main

import "fmt"

func main() {

	var year int = 2018
	fmt.Println("Year is", year)
	fmt.Printf("Type of year is %T\n", year)

	//integer wrap around
	var red uint8 = 255
	red++
	fmt.Println(red) // print 0
	var number int8 = 127
	number++
	fmt.Println(number) //print -128

	// %b represents the bit value of integer
	var num int = 3
	fmt.Printf("%08b", num)

}
