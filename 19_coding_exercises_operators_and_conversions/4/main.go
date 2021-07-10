package main

import "fmt"

/*
Coding Exercise #4

Create a Go program that computes how long does it take for the Sunlight to reach the Earth if we know that the distance from the Sun to Earth is 149.6 million km and the speed of light  is 299,792,458 m / s

The formula used is: time = distance / speed
*/
func main() {
	const distanceSunToEarthInKm = (149.6 * 1_000_000)
	const lightSpeed = 299_792_458

	const distanceSunToEarthInMeters = distanceSunToEarthInKm * 1000

	const timeOfLightTravelFromSunToEarthInSeconds = distanceSunToEarthInMeters / lightSpeed
	fmt.Printf("Sunlight takes %d minutes to reach earth", timeOfLightTravelFromSunToEarthInSeconds/60)
}
