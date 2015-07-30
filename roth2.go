/**************************
** Mission: Time to Jump **
***************************
**
** That's weird, looking at this drop sequence, it looks like
** you would miss the island...
** You know that you're cruising at 3000m and that the island is 1000m
** directly northwest from the jump site, with winds blowing with you
** You decide to fix the code to ensure you'll land safely.
**
**/

package main

import (
	"fmt"
)

func main() {

	elevation := 4000
	destinationNorth := 1000
	destinationWest := 1000

	windNorth := -4
	windWest := -4

	deployChuteAtElevation := 2000
	landingLocation := simDrop(elevation,
		destinationNorth,
		destinationWest,
		windNorth,
		windWest,
		deployChuteAtElevation)

	fmt.Println("Initiate deploy at:",
		deployChuteAtElevation,
		"meters -", landingLocation)
}

func simDrop(elevation int,
	dest_north int,
	dest_west int,
	windNorth int,
	windWest int,
	deployChuteAtElevation int) string {

	locationNorth := 0
	locationWest := 0

	deployed := false
	for elevation > 0 {
		if elevation == deployChuteAtElevation {
			deployed = true
		}
		if !deployed {
			//fmt.Println("e04")
			locationNorth -= windNorth / 4
			locationWest -= windWest / 4
		} else {
			//fmt.Println("e05")
			locationNorth -= windNorth
			locationWest -= windWest
		}
		elevation -= 10
	}
	landingLocation := "You missed the target!"
	fmt.Println(dest_north, dest_west)
	fmt.Println(locationNorth, locationWest)
	if locationNorth == dest_north && locationWest == dest_west {
		landingLocation = "You hit the target!"
	}

	return landingLocation
}
