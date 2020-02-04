package main

import (
	"fmt"
	"math"
	"strconv"
)

type Location struct {
	X, Y, Z float64
}

func (loc *Location) Distance(other Location) float64 {

	dx := loc.X - other.X
	dy := loc.Y - other.Y
	dz := loc.Z - other.Z

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (loc *Location) ToString() string {

	pointsArr := make([]Location, 10)

	res := ""

	for i := 0; i < len(pointsArr); i++ {
		res += pointsArr[i].ToString()
	}

	return "[" + strconv.FormatFloat(loc.X, 'f', -1, 64) + ", " +
		strconv.FormatFloat(loc.Y, 'f', -1, 64) + ", " +
		strconv.FormatFloat(loc.Z, 'f', -1, 64) + "]"

}

func main() {

	var p1 = Location{10.0, 11.0, 12.0}
	var p2 = Location{30.0, 40.0, 50.0}

	fmt.Printf("p1 %s\n", p1.ToString())
	fmt.Printf("p2 %s\n", p2.ToString())

	fmt.Printf("Distance between points: %f\n", p1.Distance(p2))

	fmt.Println("main done...")
}
