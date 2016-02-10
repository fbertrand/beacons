// postion.go
package beacons

import "math"

type Position struct {
	x, y int
}

func MakePosition(x_init, y_init int) Position {
	return Position{x: x_init, y: y_init}
}

func comp_distance(p1, p2 Position) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) +
		math.Pow(float64(p1.y-p2.y), 2))
}
