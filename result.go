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

func NewResult(resultId int, resultOpponentId int, resultGoals int, resultOpponentGoals int, resultGoalsAtHalfTime int, resultOpponentGoalsAtHalfTime int, resultIsHomeGame bool, resultSeasonId int, resultDate time.Time) Result {
	return Result{resultId, resultOpponentId, resultGoals, resultOpponentGoals, resultGoalsAtHalfTime, resultOpponentGoalsAtHalfTime, resultIsHomeGame, resultSeasonId, resultDate, 0}
}

func NewResultWithRound(resultId int, resultOpponentId int, resultGoals int, resultOpponentGoals int, resultGoalsAtHalfTime int, resultOpponentGoalsAtHalfTime int, resultIsHomeGame bool, resultSeasonId int, resultDate time.Time, resultRound int) Result {
	return Result{resultId, resultOpponentId, resultGoals, resultOpponentGoals, resultGoalsAtHalfTime, resultOpponentGoalsAtHalfTime, resultIsHomeGame, resultSeasonId, resultDate, resultRound}
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

func scoredInSecondHalf(r Result) bool {
	return r.GoalsAtHalfTime < r.Goals
}

func concededInFirstHalf(r Result) bool {
	return r.OpponentGoalsAtHalfTime > 0
}

func concededInSecondHalf(r Result) bool {
	return r.OpponentGoalsAtHalfTime < r.OpponentGoals
}