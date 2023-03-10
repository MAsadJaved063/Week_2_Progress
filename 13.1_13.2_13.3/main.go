package main

import "fmt"

type celsius float64
type kelvin float64

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}
func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Print(k, "K is ", c, "C")

	var f celsius = 127.0
	j := celsiusToKelvin(f)
	fmt.Print(c, "\nC is ", j, "K")
}
