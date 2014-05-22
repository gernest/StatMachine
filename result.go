package statmachine

import "time"

type Result struct {
	id                      int
	opponentId              int
	Goals                   int
	opponentGoals           int
	goalsAtHalfTime         int
	opponentGoalsAtHalfTime int
	isHomeGame              bool
	seasonId                int
	date                    time.Time
	round                   int
}

func NewResult(resultId int, resultOpponentId int, resultGoals int, resultOpponentGoals int, resultGoalsAtHalfTime int, resultOpponentGoalsAtHalfTime int, resultIsHomeGame bool, resultSeasonId int, resultDate time.Time) Result {
	return Result{resultId, resultOpponentId, resultGoals, resultOpponentGoals, resultGoalsAtHalfTime, resultOpponentGoalsAtHalfTime, resultIsHomeGame, resultSeasonId, resultDate, 0}
}

func NewResultWithRound(resultId int, resultOpponentId int, resultGoals int, resultOpponentGoals int, resultGoalsAtHalfTime int, resultOpponentGoalsAtHalfTime int, resultIsHomeGame bool, resultSeasonId int, resultDate time.Time, resultRound int) Result {
	return Result{resultId, resultOpponentId, resultGoals, resultOpponentGoals, resultGoalsAtHalfTime, resultOpponentGoalsAtHalfTime, resultIsHomeGame, resultSeasonId, resultDate, resultRound}
}

func goalsScored(r Result) int {
	return r.Goals
}

func goalsConceded(r Result) int {
	return r.opponentGoals
}
func isWin(r Result) bool {
	return r.Goals > r.opponentGoals
}

func isLoss(r Result) bool {
	return r.Goals < r.opponentGoals
}

func isDraw(r Result) bool {
	return r.Goals == r.opponentGoals
}

func scoredAGoal(r Result) bool {
	return r.Goals > 0
}

func concededAGoal(r Result) bool {
	return r.opponentGoals > 0
}

func scoredInFirstHalf(r Result) bool {
	return r.goalsAtHalfTime > 0
}

func scoredInSecondHalf(r Result) bool {
	return r.goalsAtHalfTime < r.Goals
}

func concededInFirstHalf(r Result) bool {
	return r.opponentGoalsAtHalfTime > 0
}

func concededInSecondHalf(r Result) bool {
	return r.opponentGoalsAtHalfTime < r.opponentGoals
}

func (r Result) OpponentGoals() int {
	return r.opponentGoals
}

func (r Result) IsHomeGame() bool {
	return r.isHomeGame
}

func (r Result) Date() time.Time {
	return r.date
}

func (r Result) Round() int {
	return r.round
}

func (r Result) SeasonId() int {
	return r.seasonId
}
