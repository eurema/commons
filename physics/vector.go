package physics

import (
	"encoding/json"
	"errors"
	"math"
)

type Vector struct {
	x float64
	y float64
}

func NewVector(from Point, to Point) (*Vector, error) {
	v := new(Vector)
	v.x = float64(to.PosX) - float64(from.PosX)
	v.y = float64(to.PosY) - float64(from.PosY)
	if err := v.isValidCoords(v.x, v.y); err != nil {
		return nil, err
	}
	return v, nil
}

func (v Vector) Copy() *Vector {
	nv := new(Vector)
	nv.x = v.x
	nv.y = v.y
	return nv
}

func (v Vector) Perpendicular() *Vector {
	nv := new(Vector)
	nv.x = v.y
	nv.y = -v.x
	return nv
}

func (v *Vector) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"x":   v.x,
		"y":   v.y,
		"ang": v.AngleDegrees(),
	})
}

func (v *Vector) UnmarshalJSON(b []byte) error {
	var tmp struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	v.x = tmp.X
	v.y = tmp.Y
	if err := v.isValidCoords(v.x, v.y); err != nil {
		return err
	}
	return nil
}

func (v *Vector) Normalize() *Vector {
	length := v.Length()
	if length > 0 {
		v.Scale(100 / length)
	}
	return v
}

func (v *Vector) SetLength(length float64) (*Vector, error) {
	if length == 0 {
		return nil, errors.New("vector can not have zero length")
	}
	v.Scale(length / v.Length())
	return v, nil
}

func (v *Vector) SetX(x float64) (*Vector, error) {
	if err := v.isValidCoords(x, v.y); err != nil {
		return nil, err
	}
	v.x = x
	return v, nil
}

func (v *Vector) SetY(y float64) (*Vector, error) {
	if err := v.isValidCoords(v.x, y); err != nil {
		return nil, err
	}
	v.y = y
	return v, nil
}

func (v *Vector) Invert() *Vector {
	v.x = -v.x
	v.y = -v.y
	return v
}

func (v *Vector) Scale(t float64) (*Vector, error) {
	if t == 0 {
		return nil, errors.New("vector can not have zero length")
	}
	v.x *= t
	v.y *= t
	return v, nil
}

func (v *Vector) Sin() float64 {
	return v.y / v.Length()
}

func (v *Vector) Cos() float64 {
	return v.x / v.Length()
}

// Angle returns the angle of the vector with the X axis
func (v *Vector) Angle() float64 {
	return math.Atan2(v.y, v.x)
}

func (v *Vector) AngleDegrees() float64 {
	return v.Angle() * 180 / math.Pi
}

func (v *Vector) OppositeAngle() float64 {
	return math.Acos(v.Cos())
}

func (v *Vector) AddAngleDegree(degree float64) *Vector {
	newAngle := v.AngleDegrees() + degree
	newAngle *= math.Pi / 180

	length := v.Length()
	v.x = length * math.Cos(newAngle)
	v.y = length * math.Sin(newAngle)
	return v
}

func (v *Vector) Length() float64 {
	return math.Hypot(v.x, v.y)
}

func (v *Vector) Add(vector *Vector) (*Vector, error) {
	x := v.x + vector.x
	y := v.y + vector.y
	if err := v.isValidCoords(x, y); err != nil {
		return nil, err
	}
	v.x = x
	v.y = y
	return v, nil
}

func (v *Vector) Sub(vector *Vector) (*Vector, error) {
	x := v.x - vector.x
	y := v.y - vector.y
	if err := v.isValidCoords(x, y); err != nil {
		return nil, err
	}
	v.x = x
	v.y = y
	return v, nil
}

func (v *Vector) TargetFrom(point Point) Point {
	return Point{
		point.PosX + int(math.Round(v.x)),
		point.PosY + int(math.Round(v.y)),
	}
}

func (v *Vector) GetX() float64 {
	return v.x
}

func (v *Vector) GetY() float64 {
	return v.y
}

func (v *Vector) IsEqualTo(b *Vector) bool {
	return b.y == v.y && b.x == v.x
}

func (v *Vector) AngleWith(b *Vector) float64 {
	//http://onlinemschool.com/math/assistance/vector/angl/
	copyMe := v.Copy().Normalize()
	copyOther := b.Copy().Normalize()

	dotProduct := (copyMe.x * copyOther.x) + (copyMe.y * copyOther.y)
	cos := dotProduct / (copyMe.Length() * copyOther.Length())
	ang := math.Round(math.Acos(cos)*(180/math.Pi)*100) / 100
	if copyMe.y > copyOther.y {
		ang *= -1
	}
	return ang
}

func (v *Vector) IsObstacle(from Point, obstacle Point) bool {
	to := v.TargetFrom(from)
	a := from.DistanceTo(obstacle)
	b := obstacle.DistanceTo(to)
	hypo := from.DistanceTo(to)
	return math.Round(a+b-hypo) < 0.1
}

func (v *Vector) isValidCoords(x, y float64) error {
	if x == 0 && y == 0 {
		return errors.New("vector can not have zero length")
	}
	return nil
}
