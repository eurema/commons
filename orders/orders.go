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

// AskOrderData is the expected format of the data field of an order when it's type is Ask
type AskOrderData struct {
	Question     string   `json:"question"`
	Alternatives []string `json:"alternatives"`
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
	// ASK is an order to interrupt the game (only dev mode) to ask the user (through browser) what the bot should do.
	// this order may be used by the bot to create data to be used on machine learning.
	ASK OrderType = "ASK"
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

// GetKickOrderData returns the Data order field in KickOrderData format
func (o *Order) GetAskOrderData() AskOrderData {
	return o.Data.(AskOrderData)
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
	case tmp.Type == JUMP:
		var jump JumpOrderData
		err = json.Unmarshal(tmp.Data, &jump)
		o.Data = jump
	case tmp.Type == KICK:
		var mov KickOrderData
		err = json.Unmarshal(tmp.Data, &mov)
		o.Data = mov
	case tmp.Type == ASK:
		var ask AskOrderData
		err = json.Unmarshal(tmp.Data, &ask)
		o.Data = ask
	case tmp.Type == CATCH:
		o.Data = nil
	default:
		err = errors.New(fmt.Sprintf("Unknow order type %s", tmp.Type))
	}
	return err
}

func NewMoveOrder(velocity physics.Velocity) Order {
	return Order{
		Type: MOVE,
		Data: MoveOrderData{Velocity: velocity},
	}
}

func NewJumpOrder(velocity physics.Velocity) Order {
	return Order{
		Type: JUMP,
		Data: JumpOrderData{Velocity: velocity},
	}
}

func NewKickOrder(velocity physics.Velocity) Order {
	return Order{
		Type: KICK,
		Data: KickOrderData{Velocity: velocity},
	}
}

func NewAskOrder(question string, alternatives []string) Order {
	return Order{
		Type: ASK,
		Data: AskOrderData{
			Question:     question,
			Alternatives: alternatives,
		},
	}
}

func NewCatchOrder() Order {
	return Order{
		Type: CATCH,
	}
}
