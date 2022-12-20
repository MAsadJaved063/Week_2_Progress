package main

import (
	"fmt"
	"math/big"
)

func main() {
	LightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Printf("Andromeda Galaxy is %v km away.\n", distance)
	seconds := new(big.Int)
	seconds.Div(distance, LightSpeed)
	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Printf("This is %v days of travel at light speed.\n", days)

	const distance_con = 24000000000000000000
	const lightSpeed_con = 299792
	const secondsPerDay_con = 86400
	const days_con = distance_con / lightSpeed_con / secondsPerDay_con
	fmt.Println("Andromeda Galaxy is", days_con, "light days away.")
}
