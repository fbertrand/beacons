// Room.go
package beacons

import "math"

type Beacon struct {
	Id       int
	Position Position
	Power    float64
}

func (this Beacon) signal_power(distance float64) float64 {
	return this.Power / (4 * math.Pi * math.Pow(distance, 2))
}

type Room struct {
	Beacons []Beacon
}
