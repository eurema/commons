package orders

import (
	"encoding/json"
	"github.com/makeitplay/arena/physics"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createMoveOrder(from physics.Point, to physics.Point, speed float64) Order {
	p, _ := physics.NewVector(from, to)
	vec := physics.NewZeroedVelocity(*p)
	vec.Speed = speed
	return Order{
		Type: MOVE,
		Data: MoveOrderData{vec},
	}
}

func createKickOrder(from physics.Point, to physics.Point, speed float64) Order {
	p, _ := physics.NewVector(from, to)
	vec := physics.NewZeroedVelocity(*p)
	vec.Speed = speed
	return Order{
		Type: KICK,
		Data: KickOrderData{vec},
	}
}

func TestMarshalMoveOrder(t *testing.T) {
	order := createMoveOrder(physics.Point{}, physics.Point{PosX: 5, PosY: -14}, 50)
	cont, err := json.Marshal(order)
	if err != nil {
		t.Errorf("Fail on marshal order: %s", err.Error())
	} else {
		expected := "{\"order\":\"MOVE\",\"data\":{\"velocity\":{\"direction\":{\"ang\":-70.3461759419467,\"x\":5,\"y\":-14},\"speed\":50}}}"
		assert.Equal(t, expected, string(cont))
	}
}

func TestUnmarshalMoveOrder(t *testing.T) {
	input := []byte("{\"order\":\"MOVE\",\"data\":{\"velocity\":{\"direction\":{\"ang\":-70.3461759419467,\"x\":5,\"y\":-14},\"speed\":50}}}")
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
	order := createKickOrder(physics.Point{}, physics.Point{PosX: 5, PosY: -14}, 50)
	cont, err := json.Marshal(order)
	if err != nil {
		t.Errorf("Fail on marshal order: %s", err.Error())
	} else {
		expected := "{\"order\":\"KICK\",\"data\":{\"velocity\":{\"direction\":{\"ang\":-70.3461759419467,\"x\":5,\"y\":-14},\"speed\":50}}}"
		assert.Equal(t, expected, string(cont))
	}
}

func TestUnmarshalKickOrder(t *testing.T) {
	input := []byte("{\"order\":\"KICK\",\"data\":{\"velocity\":{\"direction\":{\"ang\":-70.3461759419467,\"x\":5,\"y\":-14},\"speed\":50}}}")
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
