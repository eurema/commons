package physics

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestVector_Add(t *testing.T) {

	vectorA, _ := NewVector(Point{0, 0}, Point{1, 0})
	vectorB, _ := NewVector(Point{0, 0}, Point{1, 0})
	vectorC, _ := NewVector(Point{0, 0}, Point{-1, 0})
	vectorD, _ := NewVector(Point{0, 0}, Point{0, 1})

	velA := NewZeroedVelocity(*vectorA.Normalize())
	velB := NewZeroedVelocity(*vectorB.Normalize())
	velC := NewZeroedVelocity(*vectorC.Normalize())
	velD := NewZeroedVelocity(*vectorD.Normalize())

	velA.Speed = 100
	velB.Speed = 100
	velA.Add(velB)
	assert.Equal(t, float64(200), velA.Speed)
	assert.Equal(t, float64(100), velB.Speed)

	velA.Speed = 100
	velC.Speed = 50
	velA.Add(velC)
	assert.Equal(t, float64(50), velA.Speed)
	assert.Equal(t, float64(50), velC.Speed)

	velA.Speed = 100
	velD.Speed = 100
	velA.Add(velD)
	assert.Equal(t, math.Round(141), math.Round(velA.Speed)) // SQRT (100^2 + 100^2)
	assert.Equal(t, float64(100), velD.Speed)

}
