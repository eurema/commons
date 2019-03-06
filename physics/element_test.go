package physics

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestElement_LineCollides(t *testing.T) {
	table := map[string]struct {
		a        Point
		b        Point
		elPos    Point
		elRadius int
		margin   float64

		expectedCollide bool
		expectedCol1    *Point
		expectedCol2    *Point
	}{
		"Element axis 0,0 lin in X ": {
			Point{-10, 0},
			Point{10, 0},
			Point{0, 0},
			4,
			0,
			true,
			&Point{-4, 0},
			&Point{4, 0},
		},
		"Element axis 10,0 lin in X ": {
			Point{-10, 0},
			Point{10, 0},
			Point{10, 0},
			4,
			0,
			true,
			&Point{6, 0},
			&Point{14, 0},
		},
		"Element axis -10,0 lin in X ": {
			Point{-10, 0},
			Point{10, 0},
			Point{-10, 0},
			4,
			0,
			true,
			&Point{-14, 0},
			&Point{-6, 0},
		},
		"Element axis 0,5 lin in X ": {
			Point{-10, 0},
			Point{10, 0},
			Point{0, 5},
			4,
			0,
			false,
			nil,
			nil,
		},
		"Element axis 0,5 lin in 0,5>10,5 ": {
			Point{0, 5},
			Point{10, 5},
			Point{0, 5},
			4,
			0,
			true,
			&Point{-4, 5},
			&Point{4, 5},
		},
		"Element axis 0,9 lin in 0,5>10,5 ": {
			Point{0, 5},
			Point{10, 5},
			Point{0, 9},
			4,
			0,
			true,
			&Point{0, 5},
			nil,
		},
		"Element axis 0,9 lin in 0,2>10,2 ": {
			Point{0, 2},
			Point{10, 2},
			Point{0, 9},
			4,
			0,
			false,
			nil,
			nil,
		},
		"Diagonal": {
			Point{-30, -40},
			Point{30, 40},
			Point{0, 0},
			2,
			6.0,
			true,
			&Point{-5, -6},
			&Point{5, 6},
		},
	}

	for title, set := range table {
		element := Element{}
		element.Size = set.elRadius * 2
		element.Coords = set.elPos

		collide, col1, col2 := element.LineCollides(set.a, set.b, set.margin)
		if set.expectedCollide {
			assert.True(t, collide, title)
		} else {
			assert.False(t, collide, title)
		}

		assert.Equal(t, set.expectedCol1, col1, title)
		assert.Equal(t, set.expectedCol2, col2, title)
	}
}

func TestElement_VectorCollides(t *testing.T) {
	element := Element{}
	element.Size = 10
	element.Coords = Point{0, 0}
	//radios: 7.07
	var collisionPoint *Point

	vecA := NewVector(Point{}, Point{12, 0})

	// no collision horizontal
	assert.Nil(t, element.VectorCollides(*vecA, Point{-20, 0}, 0.0))
	assert.Nil(t, element.VectorCollides(*vecA, Point{13, 0}, 0.0))
	assert.Nil(t, element.VectorCollides(*vecA, Point{0, 11}, 0.0))
	assert.Nil(t, element.VectorCollides(*vecA, Point{0, -5}, 0.0))

	//very close at the vector end
	assert.Nil(t, element.VectorCollides(*vecA, Point{-17, 0}, 0.0))
	//very close at the vector beginning
	assert.Nil(t, element.VectorCollides(*vecA, Point{5, 0}, 0.0))
	//

	collisionPoint = element.VectorCollides(*vecA, Point{-11, 0}, 0.0)
	assert.Equal(t, &Point{-5, 0}, collisionPoint)

	collisionPoint = element.VectorCollides(*vecA, Point{-11, 0}, 5)
	assert.Equal(t, &Point{-10, 0}, collisionPoint)

	collisionPoint = element.VectorCollides(*vecA, Point{0, 0}, 0.0)
	assert.Equal(t, &Point{5, 0}, collisionPoint)

	collisionPoint = element.VectorCollides(*vecA, Point{-5, 3}, 0.0)
	assert.Equal(t, &Point{-4, 3}, collisionPoint)

	//
	element.Coords = Point{20, 20}

	vecB := NewVector(Point{}, Point{10, 10})
	vecB.SetLength(50)

	collisionPoint = element.VectorCollides(*vecB, Point{0, 0}, 0.0)
	assert.Equal(t, &Point{16, 16}, collisionPoint)
}

func TestElement_HasCollided(t *testing.T) {
	elementA := Element{
		Size:     9,
		Coords:   Point{},
		Velocity: NewZeroedVelocity(*NewVector(Point{}, Point{1, 0}).Normalize()),
	}
	elementB := Element{
		Size:     12,
		Coords:   Point{PosX: 50},
		Velocity: NewZeroedVelocity(*NewVector(Point{}, Point{1, 0}).Normalize())}

	elementsBodySpace := float64(elementB.Size+elementA.Size) / 2

	input := []struct {
		Name    string
		A       Point
		B       Point
		Collide bool
		Dist    float64
	}{
		// X
		{"X1", Point{}, Point{50, 0}, false, float64(50) - elementsBodySpace},
		{"X2", Point{}, Point{25, 0}, false, float64(25) - elementsBodySpace},
		{"X3", Point{}, Point{11, 0}, false, float64(11) - elementsBodySpace},
		{"X4", Point{}, Point{10, 0}, true, float64(10) - elementsBodySpace},
		{"X5", Point{}, Point{0, 0}, true, float64(0) - elementsBodySpace},
		{"X6", Point{}, Point{-10, 0}, true, float64(10) - elementsBodySpace},
		{"X7", Point{}, Point{-11, 0}, false, float64(11) - elementsBodySpace},
		{"X8", Point{}, Point{-25, 0}, false, float64(25) - elementsBodySpace},

		// Y
		{"Y1", Point{}, Point{0, 50}, false, float64(50) - elementsBodySpace},
		{"Y2", Point{}, Point{0, 25}, false, float64(25) - elementsBodySpace},
		{"Y3", Point{}, Point{0, 11}, false, float64(11) - elementsBodySpace},
		{"Y4", Point{}, Point{0, 10}, true, float64(10) - elementsBodySpace},
		{"Y5", Point{}, Point{0, 0}, true, float64(0) - elementsBodySpace},
		{"Y6", Point{}, Point{0, -10}, true, float64(10) - elementsBodySpace},
		{"Y7", Point{}, Point{0, -11}, false, float64(11) - elementsBodySpace},
		{"Y8", Point{}, Point{0, -25}, false, float64(25) - elementsBodySpace},

		// diagonal

		{"D1", Point{}, Point{2, 5}, true, float64(5.38516480713) - elementsBodySpace},
		{"D2", Point{}, Point{5, 9}, true, float64(10.295630141) - elementsBodySpace},
		{"D3", Point{}, Point{1, 1}, true, float64(1.41421356237) - elementsBodySpace},
		{"D4", Point{}, Point{0, 0}, true, float64(0) - elementsBodySpace},
		{"D5", Point{}, Point{5, 10}, false, float64(11.1803398875) - elementsBodySpace},
		{"D6", Point{}, Point{10, 5}, false, float64(11.1803398875) - elementsBodySpace},
		{"D7", Point{}, Point{20, 20}, false, float64(28.2842712475) - elementsBodySpace},

		//diagonal negative
		{"D1", Point{}, Point{-2, 5}, true, float64(5.38516480713) - elementsBodySpace},
		{"D2", Point{}, Point{-5, 9}, true, float64(10.295630141) - elementsBodySpace},
		{"D3", Point{}, Point{1, -1}, true, float64(1.41421356237) - elementsBodySpace},
		{"D4", Point{}, Point{0, 0}, true, float64(0) - elementsBodySpace},
		{"D5", Point{}, Point{5, -10}, false, float64(11.1803398875) - elementsBodySpace},
		{"D6", Point{}, Point{-10, 5}, false, float64(11.1803398875) - elementsBodySpace},
		{"D7", Point{}, Point{20, -20}, false, float64(28.2842712475) - elementsBodySpace},
		{"D7", Point{}, Point{-20, -20}, false, float64(28.2842712475) - elementsBodySpace},
	}

	for _, testCase := range input {
		elementA.Coords = testCase.A
		elementB.Coords = testCase.B
		collide, dist := elementA.HasCollided(&elementB)
		if collide != testCase.Collide {
			if testCase.Collide {
				t.Errorf("%s: A and B should had collide. Error dist: %f", testCase.Name, dist)
			} else {
				t.Errorf("%s: A and B should not being colliding. Error dist: %f", testCase.Name, dist)
			}
		}

		if math.Abs(math.Abs(dist)-math.Abs(testCase.Dist)) > 0.00001 {
			t.Errorf("%s: Wrong distance dist: %f (expected %f)", testCase.Name, dist, testCase.Dist)
		}

	}

	// invert
	for _, testCase := range input {
		elementA.Coords = testCase.B
		elementB.Coords = testCase.A
		collide, dist := elementA.HasCollided(&elementB)
		if collide != testCase.Collide {
			if testCase.Collide {
				t.Errorf("%s: A and B should had collide. Error dist: %f", testCase.Name, dist)
			} else {
				t.Errorf("%s: A and B should not being colliding. Error dist: %f", testCase.Name, dist)
			}
		}

		if math.Abs(math.Abs(dist)-math.Abs(testCase.Dist)) > 0.00001 {
			t.Errorf("%s: Wrong distance dist: %f (expected %f)", testCase.Name, dist, testCase.Dist)
		}

	}
}
