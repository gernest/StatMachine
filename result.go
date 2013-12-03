package statmachine

type Result struct{
	id int
	goals int
	opponentGoals int
}

func isWin(r Result) bool{
	return r.goals>r.opponentGoals
}

func isLose(r Result) bool{
	return r.goals<r.opponentGoals
}

func isDraw(r Result) bool{
	return r.goals == r.opponentGoals
}