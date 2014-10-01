package statmachine


func FindMatchingFixtures(sourceResults []Result, targetResults []Result) []Result{
	matchingResults := []Result{}

	for _,sourceRes := range sourceResults {
		for _,targetRes := range targetResults {
			if sourceRes.OpponentId == targetRes.OpponentId && sourceRes.IsHomeGame == targetRes.IsHomeGame {
				matchingResults = append(matchingResults, targetRes)
			}
		}
	}
	return matchingResults
}