package physics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeterminant_SameLine(t *testing.T) {
	A1 := Point{}
	A2 := Point{PosX: 10}

	B1 := Point{}
	B2 := Point{PosX: 10}

	_, _, err := Determinant(A1, A2, B1, B2)
	assert.NotNil(t, err)
}

func TestDeterminant_SameLineNotOverlapping(t *testing.T) {
	A1 := Point{}
	A2 := Point{PosX: 10}

	B1 := Point{PosX: 12}
	B2 := Point{PosX: 20}

	_, _, err := Determinant(A1, A2, B1, B2)
	assert.NotNil(t, err)
}

func TestDeterminant_ParallelLines(t *testing.T) {
	A1 := Point{}
	A2 := Point{PosX: 10}

	B1 := Point{PosY: 10, PosX: 12}
	B2 := Point{PosY: 10, PosX: 20}

	_, _, err := Determinant(A1, A2, B1, B2)
	assert.NotNil(t, err)
}

func TestDeterminant_CrossAndTouch(t *testing.T) {
	A1 := Point{}
	A2 := Point{PosX: 10}

	B1 := Point{PosX: 5, PosY: -5}
	B2 := Point{PosX: 5, PosY: 5}

	p, touch, err := Determinant(A1, A2, B1, B2)
	assert.Nil(t, err)
	assert.Equal(t, Point{PosX: 5, PosY: 0}, p)
	assert.True(t, touch)
}

func TestDeterminant_CrossAndDoNotTouch(t *testing.T) {
	A1 := Point{}
	A2 := Point{PosX: 10}

	B1 := Point{PosX: 5, PosY: 15}
	B2 := Point{PosX: 5, PosY: 5}

	p, touch, err := Determinant(A1, A2, B1, B2)
	assert.Nil(t, err)
	assert.Equal(t, Point{PosX: 5, PosY: 0}, p)
	assert.False(t, touch)
}

func TestDeterminant_DiagonalTouching(t *testing.T) {
	A1 := Point{}
	A2 := Point{PosX: 10, PosY: 10}

	B1 := Point{PosX: 0, PosY: 10}
	B2 := Point{PosX: 10, PosY: 0}

	p, touch, err := Determinant(A1, A2, B1, B2)
	assert.Nil(t, err)
	assert.Equal(t, Point{PosX: 5, PosY: 5}, p)
	assert.True(t, touch)
}

func TestDeterminant_DiagonalNotTouching(t *testing.T) {
	A1 := Point{}
	A2 := Point{PosX: 10, PosY: 10}

	B1 := Point{PosX: 0, PosY: 30}
	B2 := Point{PosX: 10, PosY: 20}

	p, touch, err := Determinant(A1, A2, B1, B2)
	assert.Nil(t, err)
	assert.Equal(t, Point{PosX: 15, PosY: 15}, p)
	assert.False(t, touch)
}
