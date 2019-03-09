package orders

import (
	"encoding/json"
	"fmt"
	"github.com/makeitplay/arena/physics"
	"github.com/pkg/errors"
)

// Order is a orders sent by the player to the game server during the LISTENING state
type Order struct {
	Type OrderType   `json:"order"`
	Data interface{} `json:"data"`
}

// MoveOrderData is the expected format of the data field of an order when it's type is MOVE
type MoveOrderData struct {
	Velocity physics.Velocity `json:"velocity"`
}

// KickOrderData is the expected format of the data field of an order when it's type is KICK
type KickOrderData struct {
	Velocity physics.Velocity `json:"velocity"`
}

// JumpOrderData is the expected format of the data field of an order when it's type is Jump
type JumpOrderData struct {
	Velocity physics.Velocity `json:"velocity"`
}

const (
	// orders sent by the PLAYER

	// MOVE is order to change the direction and speed of the player
	MOVE OrderType = "MOVE"
	// KICK is the order sent by the ball holder to release the ball and changes its direction and speed
	// the current ball direction will be summed with the new direction set by the order
	KICK OrderType = "KICK"
	// CATCH is an order to try to catch the ball, that has to being touched by the player
	CATCH OrderType = "CATCH"
	// JUMP is an action executed only by goal keepers! It allow the goal keeper to use extra speed during a short interval
	JUMP OrderType = "JUMP"
)

// GetMoveOrderData returns the Data order field in MoveOrderData format
func (o *Order) GetMoveOrderData() MoveOrderData {
	return o.Data.(MoveOrderData)
}

// GetJumpOrderData returns the Data order field in JumpOrderData format
func (o *Order) GetJumpOrderData() JumpOrderData {
	return o.Data.(JumpOrderData)
}

// GetKickOrderData returns the Data order field in KickOrderData format
func (o *Order) GetKickOrderData() KickOrderData {
	return o.Data.(KickOrderData)
}

// UnmarshalJSON implements the UnmarshalJSON interface for orders
func (o *Order) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Type OrderType       `json:"order"`
		Data json.RawMessage `json:"data"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	o.Type = tmp.Type
	switch {
	case tmp.Type == MOVE:
		var mov MoveOrderData
		err = json.Unmarshal(tmp.Data, &mov)
		o.Data = mov
	case tmp.Type == KICK:
		var mov KickOrderData
		err = json.Unmarshal(tmp.Data, &mov)
		o.Data = mov
	case tmp.Type == CATCH:
		o.Data = nil
	default:
		err = errors.New(fmt.Sprintf("Unknow order type %s", tmp.Type))
	}
	return err
}

func NewMoveOrder(veloticy physics.Velocity) Order {
	return Order{
		Type: MOVE,
		Data: MoveOrderData{Velocity: veloticy},
	}
}

func NewJumpOrder(veloticy physics.Velocity) Order {
	return Order{
		Type: JUMP,
		Data: JumpOrderData{Velocity: veloticy},
	}
}

func NewKickOrder(veloticy physics.Velocity) Order {
	return Order{
		Type: KICK,
		Data: KickOrderData{Velocity: veloticy},
	}
}

func NewCatchOrder() Order {
	return Order{
		Type: CATCH,
	}
}
