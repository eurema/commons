package physics

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestVector_AngleWith_ZeroDegree(t *testing.T) {
	testTable := map[string]struct {
		vecA *Vector
		vecB *Vector
		ang  float64
	}{
		"Same direction East": {
			vecA: NewVector(Point{0, 0}, Point{1, 0}),
			vecB: NewVector(Point{0, 0}, Point{1, 0}),
			ang:  0.0,
		},
		"Same direction North": {
			vecA: NewVector(Point{0, 0}, Point{0, 1}),
			vecB: NewVector(Point{0, 0}, Point{0, 1}),
			ang:  0.0,
		},
		"Same direction Southweast": {
			vecA: NewVector(Point{0, 0}, Point{-5, -10}),
			vecB: NewVector(Point{0, 0}, Point{-5, -10}),
			ang:  0.0,
		},

		"90 degree North": {
			vecA: NewVector(Point{0, 0}, Point{1, 0}),
			vecB: NewVector(Point{0, 0}, Point{0, 1}),
			ang:  90.0,
		},

		"90 degree South": {
			vecA: NewVector(Point{0, 0}, Point{1, 0}),
			vecB: NewVector(Point{0, 0}, Point{0, -1}),
			ang:  -90.0,
		},
		"180 degrees": {
			vecA: NewVector(Point{0, 0}, Point{1, 0}),
			vecB: NewVector(Point{0, 0}, Point{-1, 0}),
			ang:  180,
		},

		"45 degrees Northeast": {
			vecA: NewVector(Point{0, 0}, Point{1, 0}),
			vecB: NewVector(Point{0, 0}, Point{1, 1}),
			ang:  45,
		},

		"45 degrees Southeast": {
			vecA: NewVector(Point{0, 0}, Point{1, 0}),
			vecB: NewVector(Point{0, 0}, Point{1, -1}),
			ang:  -45,
		},

		"135 degrees Northweast": {
			vecA: NewVector(Point{0, 0}, Point{1, 0}),
			vecB: NewVector(Point{0, 0}, Point{-1, 1}),
			ang:  135,
		},

		"135 degrees Southweast": {
			vecA: NewVector(Point{0, 0}, Point{1, 0}),
			vecB: NewVector(Point{0, 0}, Point{-1, -1}),
			ang:  -135,
		},

		"90 both not zero": {
			vecA: NewVector(Point{0, 0}, Point{1, 1}),
			vecB: NewVector(Point{0, 0}, Point{-1, 1}),
			ang:  90,
		},
	}

	for title, conditions := range testTable {
		actualAng := conditions.vecA.AngleWith(conditions.vecB)
		assert.Equal(t, conditions.ang, actualAng, title)
	}

}

func TestVector_AddAngle(t *testing.T) {
	vecA := NewVector(Point{0, 0}, Point{100, 0})

	vecA.AddAngleDegree(90)
	assert.Equal(t, float64(90), math.Round(vecA.AngleDegrees()))
	assert.True(t, vecA.x <= 0.00000001)
	assert.Equal(t, float64(100), vecA.y)
	assert.Equal(t, float64(100), vecA.Length())

	vecA.AddAngleDegree(90)
	assert.Equal(t, float64(180), math.Round(vecA.AngleDegrees()))
	assert.Equal(t, float64(-100), vecA.x)
	assert.True(t, vecA.y <= 0.00000001)
	assert.Equal(t, float64(100), vecA.Length())

	vecA.AddAngleDegree(90)
	assert.Equal(t, float64(-90), math.Round(vecA.AngleDegrees()))
	assert.True(t, vecA.x <= 0.00000001)
	assert.Equal(t, float64(-100), vecA.y)
	assert.Equal(t, float64(100), vecA.Length())

	vecA.AddAngleDegree(90)
	assert.Equal(t, float64(0), math.Round(vecA.AngleDegrees()))
	assert.Equal(t, float64(100), vecA.x)
	assert.True(t, vecA.y <= 0.00000001)
	assert.Equal(t, float64(100), vecA.Length())

	vecA.AddAngleDegree(45)
	assert.Equal(t, float64(45), math.Round(vecA.AngleDegrees()))
	assert.True(t, math.Abs(vecA.y-vecA.x) <= 0.00000001)
	assert.Equal(t, float64(100), vecA.Length())

}
