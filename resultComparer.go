package statmachine


func FindMatchingFixtures(sourceResults []Result, targetResults []Result) []ResultPair{
	matchingResults := []ResultPair{}

	for _,sourceRes := range sourceResults {
		for _,targetRes := range targetResults {
			if sourceRes.OpponentId == targetRes.OpponentId && sourceRes.IsHomeGame == targetRes.IsHomeGame {
				matchingResults = append(matchingResults, ResultPair{sourceRes,targetRes})
			}
		}
	}
	return matchingResults
}