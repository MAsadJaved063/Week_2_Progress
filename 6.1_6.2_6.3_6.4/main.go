package main

import (
	"fmt"
	"math"
)

func main() {

	var answer float64 = 3.21
	fmt.Println("Floating number is ", answer)

	answer1 := 9.0
	fmt.Printf("Type of answer1 variable is %T\n", answer1)
	var pi64 = math.Pi
	var pi32 float32 = math.Pi
	fmt.Printf("percision of PI in single percision %v and double percision %v", pi32, pi64)

	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%0.5f\n", third)
	fmt.Printf("%0.7f\n", third)
	/*left side of floating point value represents the Total length of floating value
	including the decimal and right side represents the precision */

	fmt.Printf("%09.3f\n", third)

	//Floating point inaccuracy
	piggybank := 0.1
	piggybank += 0.2

	//give the inaccuracy
	fmt.Println(piggybank)

	//Temperature Calculation
	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, " F\n")
	//Always use Multiplication First for more accuracy
	fmt.Print((celsius*9.0/5.0)+32, " F\n")

	//Floating point comparing
	fmt.Println(piggybank == 0.3) //Gives False

	//Another way to compare it accurately
	fmt.Println(math.Abs(piggybank-0.3) < 0.0001) //Gives True

}
