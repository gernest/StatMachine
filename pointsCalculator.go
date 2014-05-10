package statmachine

const pointsForADraw = 1
const pointsForAWin = 3

func Points(res []Result) int {
	totalPoints := 0
	for _, r := range res{
		if(isWin(r)){
			totalPoints = totalPoints + pointsForAWin
		} else if (isDraw(r)) {
			totalPoints = totalPoints + pointsForADraw
		}
	}
	return totalPoints
}