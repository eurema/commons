package GameState

// State is a game state
type State string

const (
	//WaitingTeams game state when the game server is waiting for both team's players connections
	WaitingTeams State = "waiting"
	////GetReady game state when the game server is asking the players that a new cycle will start
	//GetReady State = "get-ready"
	//Listening game state when the game server is listening the player for orders
	Listening State = "listening"
	//Playing game state when the game server is executing the orders sent during the last `listening` state
	Playing State = "playing"
	//Pause game state when the game server is paused by a debug command and waiting for the `next step` signal
	Pause State = "pause"
	//Results game state when the game server is announcing the score change
	Results State = "results"
	//Over game state when the game server is is announcing the end of the game
	Over State = "game-over"
)
