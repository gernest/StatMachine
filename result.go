package statmachine

import "time"

type Result struct{
	id int
	opponentId int
	goals int
	opponentGoals int
	goalsAtHalfTime int	
	opponentGoalsAtHalfTime int	
	isHomeGame bool
	seasonId int
	date time.Time
	Round int
}

func NewResult(resultId int, resultOpponentId int, resultGoals int,  resultOpponentGoals int, resultGoalsAtHalfTime int, resultOpponentGoalsAtHalfTime int, resultIsHomeGame bool, resultSeasonId int, resultDate time.Time) Result{
	return Result{resultId, resultOpponentId, resultGoals, resultOpponentGoals, resultGoalsAtHalfTime, resultOpponentGoalsAtHalfTime, resultIsHomeGame, resultSeasonId, resultDate, 0}
}

func NewResultWithRound(resultId int, resultOpponentId int, resultGoals int,  resultOpponentGoals int, resultGoalsAtHalfTime int, resultOpponentGoalsAtHalfTime int, resultIsHomeGame bool, resultSeasonId int, resultDate time.Time, resultRound int) Result{
	return Result{resultId, resultOpponentId, resultGoals, resultOpponentGoals, resultGoalsAtHalfTime, resultOpponentGoalsAtHalfTime, resultIsHomeGame, resultSeasonId, resultDate, resultRound}
}

func goalsScored(r Result) int{
	return r.goals
}

func goalsConceded(r Result) int{
	return r.opponentGoals
}
func isWin(r Result) bool{
	return r.goals>r.opponentGoals
}

func isLoss(r Result) bool{
	return r.goals<r.opponentGoals
}

func isDraw(r Result) bool{
	return r.goals == r.opponentGoals
}

func scoredAGoal(r Result) bool{
	return r.goals>0
}

func concededAGoal(r Result) bool{
	return r.opponentGoals>0
}

func scoredInFirstHalf(r Result) bool{
	return r.goalsAtHalfTime>0
}

func scoredInSecondHalf(r Result) bool{
	return r.goalsAtHalfTime<r.goals
}

func concededInFirstHalf(r Result) bool{
	return r.opponentGoalsAtHalfTime>0
}

func concededInSecondHalf(r Result) bool{
	return r.opponentGoalsAtHalfTime<r.opponentGoals
}