package physics

import (
	"fmt"
	"math"
)

// Point represents a exact point in the field
type Point struct {
	PosX int `json:"x"`
	PosY int `json:"y"`
}

// DistanceTo finds the distance of this point to a target point
func (p *Point) DistanceTo(target Point) (distance float64) {
	catA := float64(target.PosX) - float64(p.PosX)
	catO := float64(target.PosY) - float64(p.PosY)
	return math.Hypot(catA, catO)
}

// MiddlePointTo finds a point between this point and a target point
func (p *Point) MiddlePointTo(target Point) Point {
	x := math.Abs(float64(p.PosX - target.PosX))
	y := math.Abs(float64(p.PosY - target.PosY))

	return Point{
		PosX: int(math.Round(math.Min(float64(p.PosX), float64(target.PosX)) + x)),
		PosY: int(math.Round(math.Min(float64(p.PosY), float64(target.PosY)) + y)),
	}
}

// String returns the string representation of a point
func (p *Point) String() string {
	return fmt.Sprintf("{%d, %d}", p.PosX, p.PosY)
}
