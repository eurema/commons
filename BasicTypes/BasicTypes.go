package BasicTypes

import (
	"github.com/makeitplay/arena/Physics"
	"github.com/makeitplay/arena/units"
)

// MsgType define strings acceptable as types of game msg
type MsgType string

// OrderType identifies types of orders that are acceptable by the game server
type OrderType string

// PlayerNumber identifies values for players number
type PlayerNumber string

const (
	// ORDER is the msg sent from the player to the game server
	ORDER MsgType = "order"
	// ANNOUNCEMENT is sent from the game server to the players and to the web clients to update them with a new game state
	ANNOUNCEMENT MsgType = "announcement"
	// DEBUG is a message sent by http POST request from the web client to the game server (debug mode must be on)
	DEBUG MsgType = "debug"
	// SCORE is a message sent by the game server when the score was changed
	SCORE MsgType = "score"
	// RIP is a message sent by the game server when the game server crashes
	RIP MsgType = "rip"
	// WELCOME is a message sent by the game server to each player when the new websocket connection is accepted.
	WELCOME MsgType = "welcome"
)

// State identifies game states
type State string

// PlayerSpecifications is the object that should be present in the HTTP websocket headers connection open by the player with the game server
type PlayerSpecifications struct {
	// Number identifies the number of the player in its team
	Number PlayerNumber `json:"number"`
	// InitialCoords identifies where default initial player's position is
	InitialCoords Physics.Point `json:"initial_coords"`
	// Token should be passed as an argument to the player to ensure that the connection is being openned by the expected process
	Token string `json:"token"`
	// ProtocolVersion identifies the game server communication version the player is compatible with (e.g. 1.0)
	ProtocolVersion string `json:"protocol_version"`
}

// Goal is a set of value about a goal from a team
type Goal struct {
	// Center the is coordinate of the center of the goal
	Center Physics.Point
	// Place identifies the team of this goal (the team that should defend this goal)
	Place units.TeamPlace
	// TopPole is the coordinates of the pole with a higher Y coordinate
	TopPole Physics.Point
	// BottomPole is the coordinates of the pole  with a lower Y coordinate
	BottomPole Physics.Point
}
