package commons

import (
	"github.com/makeitplay/arena/BasicTypes"
	"github.com/makeitplay/arena/Physics"
	"github.com/makeitplay/arena/units"
)

const (
	// GoalkeeperNumber defines the goalkeeper number
	GoalkeeperNumber BasicTypes.PlayerNumber = "1"
)

// HomeTeamGoal works as a constant value to help to retrieve a Goal struct with the values of the Home team goal
var HomeTeamGoal = BasicTypes.Goal{
	Place:      units.HomeTeam,
	Center:     Physics.Point{0, units.CourtHeight / 2},
	TopPole:    Physics.Point{0, units.GoalMaxY},
	BottomPole: Physics.Point{0, units.GoalMinY},
}

// AwayTeamGoal works as a constant value to help to retrieve a Goal struct with the values of the Away team goal
var AwayTeamGoal = BasicTypes.Goal{
	Place:      units.HomeTeam,
	Center:     Physics.Point{units.CourtWidth, units.CourtHeight / 2},
	TopPole:    Physics.Point{units.CourtWidth, units.GoalMaxY},
	BottomPole: Physics.Point{units.CourtWidth, units.GoalMinY},
}

// CourtCenter works as a constant value to help to retrieve a Point struct with the values of the center of the court
var CourtCenter = Physics.Point{units.CourtWidth / 2, units.CourtHeight / 2}
