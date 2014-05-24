package statmachine

func lastGameWon(res []Result) Result {
	return findFirst(res, func(r Result) bool { return r.IsWin() })
}

func lastGameLost(res []Result) Result {
	return findFirst(res, func(r Result) bool { return r.IsLoss() })
}

func lastGameDrawn(res []Result) Result {
	return findFirst(res, func(r Result) bool { return r.IsDraw() })
}

func findFirst(res []Result, f func(Result) bool) Result {
	found := Result{}
	for _, e := range res {
		if f(e) {
			found = e
			break
		}
	}
	return found
}
