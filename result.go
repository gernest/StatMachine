package statmachine

type Result struct{
	id int
	goals int
	opponentGoals int
	goalsAtHalfTime int	
	opponentGoalsAtHalfTime int	
	isHomeGame bool
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