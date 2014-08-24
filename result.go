package statmachine

import "time"

type Result struct {
	Id                      int
	OpponentId              int
	Goals                   int
	OpponentGoals           int
	GoalsAtHalfTime         int
	OpponentGoalsAtHalfTime int
	IsHomeGame              bool
	SeasonId                int
	Date                    time.Time
	Round                   int
	Cards                   CardInfo
}

type CardInfo struct {
	HomeTeamNumberOfYellowCards uint8
	HomeTeamNumberOfRedCards    uint8
	AwayTeamNumberOfYellowCards uint8
	AwayTeamNumberOfRedCards    uint8
}

func (r Result) IsWin() bool {
	return r.Goals > r.OpponentGoals
}

func (r Result) IsLoss() bool {
	return r.Goals < r.OpponentGoals
}

func (r Result) IsDraw() bool {
	return r.Goals == r.OpponentGoals
}

func (r Result) ScoredAGoal() bool {
	return r.Goals > 0
}

func (r Result) ConcededAGoal() bool {
	return r.OpponentGoals > 0
}

func (r Result) ScoredInFirstHalf() bool {
	return r.GoalsAtHalfTime > 0
}

func (r Result) ScoredInSecondHalf() bool {
	return r.GoalsAtHalfTime < r.Goals
}

func (r Result) ConcededInFirstHalf() bool {
	return r.OpponentGoalsAtHalfTime > 0
}

func (r Result) ConcededInSecondHalf() bool {
	return r.OpponentGoalsAtHalfTime < r.OpponentGoals
}

func (r Result) Points() uint8 {
	if r.IsWin() {
		return 3
	}else if r.IsDraw() {
		return 1
	}else {
		return 0
	}
}

func (r Result) FirstHalfPoints() uint8 {
	if r.GoalsAtHalfTime > r.OpponentGoalsAtHalfTime {
		return 3
	}else if r.GoalsAtHalfTime == r.OpponentGoalsAtHalfTime {
		return 1
	}else {
		return 0
	}
}

func (r Result) SecondHalfPoints() uint8 {
	if (r.Goals - r.GoalsAtHalfTime) > (r.OpponentGoals - r.OpponentGoalsAtHalfTime) {
		return 3
	}else if (r.Goals - r.GoalsAtHalfTime) == (r.OpponentGoals - r.OpponentGoalsAtHalfTime) {
		return 1
	}else {
		return 0
	}
}

