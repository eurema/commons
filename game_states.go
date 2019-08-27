package arena

// GameState is a game state
type GameState string

const (
	//WaitingTeams game state when the game server is waiting for both team's players connections
	WaitingTeams GameState = "waiting"
	//Ready game state when the game server is ready to start
	Ready GameState = "ready"
	//Listening game state when the game server is listening the player for orders
	Listening GameState = "listening"
	//Playing game state when the game server is executing the orders sent during the last `listening` state
	Playing GameState = "playing"
	//Pause game state when the game server is paused by a debug command and waiting for the `next step` signal
	Pause GameState = "pause"
	//Results game state when the game server is announcing the score change
	Results GameState = "results"
	//Over game state when the game server is is announcing the end of the game
	Over GameState = "game-over"
)
