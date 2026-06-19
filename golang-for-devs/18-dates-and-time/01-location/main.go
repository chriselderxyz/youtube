package main

import (
	"fmt"
	"os"
	"time"
)

/////////////////////////////////////////////////////////////
// time.Location
/////////////////////////////////////////////////////////////

func main() {
	loc1 := time.Local
	loc2 := time.UTC
	loc3, err := time.LoadLocation("America/New_York")

	data, err := os.ReadFile("/usr/share/zoneinfo/America/Toronto")
	if err != nil {
		panic(err)
	}

	loc4, err := time.LoadLocationFromTZData("My Timezone", data)

	loc5 := time.FixedZone("My TZ", -5*60*60)

	fmt.Println("loc1 - time.Local: ", loc1)
	fmt.Println("loc2 - time.UTC: ", loc2)
	fmt.Println("loc3 - time.LoadLocation: ", loc3)
	fmt.Println("loc4 - time.LoadLocationFromTZData ", loc4)
	fmt.Println("loc5 - Custom offset with time.FixedZone: ", loc5)
}
