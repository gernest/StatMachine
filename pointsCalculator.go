package statmachine

const pointsForADraw = 1
const pointsForAWin = 3

func Points(res []Result) int {
	totalPoints := 0
	for _, r := range res {
		if r.IsWin() {
			totalPoints = totalPoints + pointsForAWin
		} else if r.IsDraw() {
			totalPoints = totalPoints + pointsForADraw
		}
	}
	return totalPoints
}
