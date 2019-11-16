package arena

import (
	"github.com/lugobots/arena/physics"
	"github.com/lugobots/arena/units"
)

const (
	// GoalkeeperNumber defines the goalkeeper number
	GoalkeeperNumber PlayerNumber = "1"
)

// HomeTeamGoal works as a constant value to help to retrieve a Goal struct with the values of the Home team goal
var HomeTeamGoal = Goal{
	Place:      HomeTeam,
	Center:     physics.Point{PosX: 0, PosY: units.FieldHeight / 2},
	TopPole:    physics.Point{PosX: 0, PosY: units.GoalMaxY},
	BottomPole: physics.Point{PosX: 0, PosY: units.GoalMinY},
}

// AwayTeamGoal works as a constant value to help to retrieve a Goal struct with the values of the Away team goal
var AwayTeamGoal = Goal{
	Place:      HomeTeam,
	Center:     physics.Point{PosX: units.FieldWidth, PosY: units.FieldHeight / 2},
	TopPole:    physics.Point{PosX: units.FieldWidth, PosY: units.GoalMaxY},
	BottomPole: physics.Point{PosX: units.FieldWidth, PosY: units.GoalMinY},
}

// FieldCenter works as a constant value to help to retrieve a Point struct with the values of the center of the court
var FieldCenter = physics.Point{units.FieldWidth / 2, units.FieldHeight / 2}
