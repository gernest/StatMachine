package statmachine

import "time"

type Result struct {
	Id                      int
	OpponentId              int
	Goals                   int
	OpponentGoals           int
	GoalsAtHalfTime         int
	OpponentGoalsAtHalfTime int
	IsHomeGame              bool
	SeasonId                int
	Date                    time.Time
	Round                   int
}

func (r Result) IsWin() bool {
	return r.Goals > r.OpponentGoals
}

func (r Result) IsLoss() bool {
	return r.Goals < r.OpponentGoals
}

func (r Result) IsDraw() bool {
	return r.Goals == r.OpponentGoals
}

func (r Result) ScoredAGoal() bool {
	return r.Goals > 0
}

func (r Result) ConcededAGoal() bool {
	return r.OpponentGoals > 0
}

func (r Result) ScoredInFirstHalf() bool {
	return r.GoalsAtHalfTime > 0
}

func (r Result) ScoredInSecondHalf() bool {
	return r.GoalsAtHalfTime < r.Goals
}

func (r Result) ConcededInFirstHalf() bool {
	return r.OpponentGoalsAtHalfTime > 0
}

func (r Result) ConcededInSecondHalf() bool {
	return r.OpponentGoalsAtHalfTime < r.OpponentGoals
}