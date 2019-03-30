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

//find determinant to find the point the ball crossed
//https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_points_on_each_line
func Determinant(a1, a2, b1, b2 Point) (Point, bool, error) {

	div := ((a1.PosX - a2.PosX) * (b1.PosY - b2.PosY)) - ((a1.PosY - a2.PosY) * (b1.PosX - b2.PosX))
	if div == 0 {
		return Point{}, false, fmt.Errorf("invalid points, they may be in the same line")
	}
	quoX := ((a1.PosX*a2.PosY)-(a1.PosY*a2.PosX))*(b1.PosX-b2.PosX) - ((a1.PosX - a2.PosX) * (b1.PosX*b2.PosY - b1.PosY*b2.PosX))
	quoY := ((a1.PosX*a2.PosY)-(a1.PosY*a2.PosX))*(b1.PosY-b2.PosY) - ((a1.PosY - a2.PosY) * (b1.PosX*b2.PosY - b1.PosY*b2.PosX))

	crossPoints := Point{
		PosX: quoX / div,
		PosY: quoY / div,
	}

	isInAX := isBetween(crossPoints.PosX, a1.PosX, a2.PosX)
	isInAY := isBetween(crossPoints.PosY, a1.PosY, a2.PosY)
	isInBX := isBetween(crossPoints.PosX, b1.PosX, b2.PosX)
	isInBY := isBetween(crossPoints.PosY, b1.PosY, b2.PosY)

	touchLines := isInAX && isInAY && isInBX && isInBY
	return crossPoints, touchLines, nil
}

func isBetween(target, coordA, coordB int) bool {
	targetF := float64(target)
	coordAF := float64(coordA)
	coordBF := float64(coordB)
	return targetF <= math.Max(coordAF, coordBF) && targetF >= math.Min(coordAF, coordBF)
}
