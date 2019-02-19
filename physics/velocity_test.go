package physics

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestVector_Add(t *testing.T) {
	velA := NewZeroedVelocity(*NewVector(Point{0, 0}, Point{1, 0}).Normalize())
	velB := NewZeroedVelocity(*NewVector(Point{0, 0}, Point{1, 0}).Normalize())
	velC := NewZeroedVelocity(*NewVector(Point{0, 0}, Point{-1, 0}).Normalize())
	velD := NewZeroedVelocity(*NewVector(Point{0, 0}, Point{0, 1}).Normalize())

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
