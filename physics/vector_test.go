package physics

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestVector_AngleWith_ZeroDegree(t *testing.T) {
	type tTable struct {
		vecA *Vector
		vecB *Vector
		ang  float64
	}
	testTable := map[string]tTable{}

	caseSample := tTable{}
	caseSample.vecA, _ = NewVector(Point{0, 0}, Point{1, 0})
	caseSample.vecB, _ = NewVector(Point{0, 0}, Point{1, 0})
	caseSample.ang = 0.0
	testTable["Same direction East"] = caseSample

	caseSample = tTable{}
	caseSample.vecA, _ = NewVector(Point{0, 0}, Point{0, 1})
	caseSample.vecB, _ = NewVector(Point{0, 0}, Point{0, 1})
	caseSample.ang = 0.0
	testTable["Same direction North"] = caseSample

	caseSample = tTable{}
	caseSample.vecA, _ = NewVector(Point{0, 0}, Point{-5, -10})
	caseSample.vecB, _ = NewVector(Point{0, 0}, Point{-5, -10})
	caseSample.ang = 0.0
	testTable["Same direction Southweast"] = caseSample

	caseSample = tTable{}
	caseSample.vecA, _ = NewVector(Point{0, 0}, Point{1, 0})
	caseSample.vecB, _ = NewVector(Point{0, 0}, Point{0, 1})
	caseSample.ang = 90.0
	testTable["90 degree North"] = caseSample

	caseSample = tTable{}
	caseSample.vecA, _ = NewVector(Point{0, 0}, Point{1, 0})
	caseSample.vecB, _ = NewVector(Point{0, 0}, Point{0, -1})
	caseSample.ang = -90.0
	testTable["90 degree South"] = caseSample

	caseSample = tTable{}
	caseSample.vecA, _ = NewVector(Point{0, 0}, Point{1, 0})
	caseSample.vecB, _ = NewVector(Point{0, 0}, Point{-1, 0})
	caseSample.ang = 180
	testTable["180 degrees"] = caseSample

	caseSample = tTable{}
	caseSample.vecA, _ = NewVector(Point{0, 0}, Point{1, 0})
	caseSample.vecB, _ = NewVector(Point{0, 0}, Point{1, 1})
	caseSample.ang = 45
	testTable["45 degrees Northeast"] = caseSample

	caseSample = tTable{}
	caseSample.vecA, _ = NewVector(Point{0, 0}, Point{1, 0})
	caseSample.vecB, _ = NewVector(Point{0, 0}, Point{1, -1})
	caseSample.ang = -45
	testTable["45 degrees Southeast"] = caseSample

	caseSample = tTable{}
	caseSample.vecA, _ = NewVector(Point{0, 0}, Point{1, 0})
	caseSample.vecB, _ = NewVector(Point{0, 0}, Point{-1, 1})
	caseSample.ang = 135
	testTable["135 degrees Northweast"] = caseSample

	caseSample = tTable{}
	caseSample.vecA, _ = NewVector(Point{0, 0}, Point{1, 0})
	caseSample.vecB, _ = NewVector(Point{0, 0}, Point{-1, -1})
	caseSample.ang = -135
	testTable["135 degrees Southweast"] = caseSample

	caseSample = tTable{}
	caseSample.vecA, _ = NewVector(Point{0, 0}, Point{1, 1})
	caseSample.vecB, _ = NewVector(Point{0, 0}, Point{-1, 1})
	caseSample.ang = 90
	testTable["90 both not zero"] = caseSample

	for title, conditions := range testTable {
		actualAng := conditions.vecA.AngleWith(conditions.vecB)
		assert.Equal(t, conditions.ang, actualAng, title)
	}

}

func TestVector_AddAngle(t *testing.T) {
	vecA, _ := NewVector(Point{0, 0}, Point{100, 0})

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
