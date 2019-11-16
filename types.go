package arena

import (
	"github.com/lugobots/arena/physics"
)

// TeamPlace defines a side of the team during the game (left for home team, and right for the away team)
type TeamPlace string

// HomeTeam identify the home team
const HomeTeam TeamPlace = "home"

// AwayTeam identify the home team
const AwayTeam TeamPlace = "away"

// MsgType define strings acceptable as types of game msg
type MsgType string

// PlayerNumber identifies values for players number
type PlayerNumber string

// PlayerSpecifications is the object that should be present in the HTTP websocket headers connection open by the player with the game server
type PlayerSpecifications struct {
	// Number identifies the number of the player in its team
	Number PlayerNumber `json:"number"`
	// InitialCoords identifies where default initial player's position is
	InitialCoords physics.Point `json:"initial_coords"`
	// Token should be passed as an argument to the player to ensure that the connection is being openned by the expected process
	Token string `json:"token"`
	// ProtocolVersion identifies the game server communication version the player is compatible with (e.g. 1.0)
	ProtocolVersion string `json:"protocol_version"`
}

// Goal is a set of value about a goal from a team
type Goal struct {
	// Center the is coordinate of the center of the goal
	Center physics.Point
	// Place identifies the team of this goal (the team that should defend this goal)
	Place TeamPlace
	// TopPole is the coordinates of the pole with a higher Y coordinate
	TopPole physics.Point
	// BottomPole is the coordinates of the pole  with a lower Y coordinate
	BottomPole physics.Point
}
