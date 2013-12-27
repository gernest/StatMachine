package statmachine

type Result struct{
	id int
	goals int
	opponentGoals int
	goalsAtHalfTime int	
	opponentGoalsAtHalfTime int	
	isHomeGame bool
}

func NewResult(resultId int, resultGoals int, resultOpponentGoals int, resultGoalsAtHalfTime int, resultOpponentGoalsAtHalfTime int, resultIsHomeGame bool) Result{
	return Result{resultId, resultGoals, resultOpponentGoals, resultGoalsAtHalfTime, resultOpponentGoalsAtHalfTime, resultIsHomeGame}
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