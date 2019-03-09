package orders

import "github.com/makeitplay/arena"

// OrderType identifies types of orders that are acceptable by the game server
type OrderType string

const (
	// ORDER is the msg sent from the player to the game server
	ORDER arena.MsgType = "order"
	// ANNOUNCEMENT is sent from the game server to the players and to the web clients to update them with a new game state
	ANNOUNCEMENT arena.MsgType = "announcement"
	// DEBUG is a message sent by http POST request from the web client to the game server (debug mode must be on)
	DEBUG arena.MsgType = "debug"
	// SCORE is a message sent by the game server when the score was changed
	SCORE arena.MsgType = "score"
	// RIP is a message sent by the game server when the game server crashes
	RIP arena.MsgType = "rip"
	// WELCOME is a message sent by the game server to each player when the new websocket connection is accepted.
	WELCOME arena.MsgType = "welcome"
)
