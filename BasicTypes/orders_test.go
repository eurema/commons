package BasicTypes

import (
	"encoding/json"
	"github.com/makeitplay/arena/Physics"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createMoveOrder(from Physics.Point, to Physics.Point, speed float64) Order {
	vec := Physics.NewZeroedVelocity(*Physics.NewVector(from, to))
	vec.Speed = speed
	return Order{
		Type: MOVE,
		Data: MoveOrderData{vec},
	}
}

func createKickOrder(from Physics.Point, to Physics.Point, speed float64) Order {
	vec := Physics.NewZeroedVelocity(*Physics.NewVector(from, to))
	vec.Speed = speed
	return Order{
		Type: KICK,
		Data: KickOrderData{vec},
	}
}

func TestMarshalMoveOrder(t *testing.T) {
	order := createMoveOrder(Physics.Point{}, Physics.Point{PosX: 5, PosY: -14}, 50)
	cont, err := json.Marshal(order)
	if err != nil {
		t.Errorf("Fail on marshal order: %s", err.Error())
	} else {
		expected := "{\"order\":\"MOVE\",\"data\":{\"velocity\":{\"direction\":{\"x\":5,\"y\":-14},\"speed\":50}}}"
		assert.Equal(t, expected, string(cont))
	}
}

func TestUnmarshalMoveOrder(t *testing.T) {
	input := []byte("{\"order\":\"MOVE\",\"data\":{\"velocity\":{\"direction\":{\"x\":5,\"y\":-14},\"speed\":50}}}")
	var order Order
	err := json.Unmarshal(input, &order)
	if err != nil {
		t.Errorf("Fail on unmarshal order: %s", err.Error())
	} else {
		assert.Equal(t, order.Type, MOVE)
		moveOrder := order.GetMoveOrderData()
		assert.Equal(t, float64(50), moveOrder.Velocity.Speed)
		assert.Equal(t, float64(5.0), moveOrder.Velocity.Direction.GetX())
		assert.Equal(t, float64(-14), moveOrder.Velocity.Direction.GetY())
	}
}

func TestMarshalKickOrder(t *testing.T) {
	order := createKickOrder(Physics.Point{}, Physics.Point{PosX: 5, PosY: -14}, 50)
	cont, err := json.Marshal(order)
	if err != nil {
		t.Errorf("Fail on marshal order: %s", err.Error())
	} else {
		expected := "{\"order\":\"KICK\",\"data\":{\"velocity\":{\"direction\":{\"x\":5,\"y\":-14},\"speed\":50}}}"
		assert.Equal(t, expected, string(cont))
	}
}

func TestUnmarshalKickOrder(t *testing.T) {
	input := []byte("{\"order\":\"KICK\",\"data\":{\"velocity\":{\"direction\":{\"x\":5,\"y\":-14},\"speed\":50}}}")
	var order Order
	err := json.Unmarshal(input, &order)
	if err != nil {
		t.Errorf("Fail on unmarshal order: %s", err.Error())
	} else {
		assert.Equal(t, order.Type, KICK)
		kickOrder := order.GetKickOrderData()
		assert.Equal(t, float64(50), kickOrder.Velocity.Speed)
		assert.Equal(t, float64(5.0), kickOrder.Velocity.Direction.GetX())
		assert.Equal(t, float64(-14), kickOrder.Velocity.Direction.GetY())
	}
}
