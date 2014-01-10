package statmachine

type Result struct{
	id int
	goals int
	opponentGoals int
	goalsAtHalfTime int	
	opponentGoalsAtHalfTime int	
	isHomeGame bool
	seasonId int
}

func NewResult(resultId int, resultGoals int, resultOpponentGoals int, resultGoalsAtHalfTime int, resultOpponentGoalsAtHalfTime int, resultIsHomeGame bool, resultSeasonId int) Result{
	return Result{resultId, resultGoals, resultOpponentGoals, resultGoalsAtHalfTime, resultOpponentGoalsAtHalfTime, resultIsHomeGame, resultSeasonId}
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