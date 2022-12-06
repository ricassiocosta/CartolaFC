package entity

import (
	"strconv"
	"time"
)

type MatchResult struct {
	teamAScore int
	teamBScore int
}

func NewMatchResult(teamAscore, teamBScore int) *MatchResult {
	return &MatchResult{
		teamAScore: teamAscore,
		teamBScore: teamBScore,
	}
}

func (m *MatchResult) GetResult() string {
	return strconv.Itoa(m.teamAScore) + "-" + strconv.Itoa(m.teamBScore)
}

type Match struct {
	ID      string
	TeamA   *Team
	TeamB   *Team
	Date    time.Time
	Status  string
	Result  MatchResult
	Actions []GameAction
}

func NewMatch(id string, teamA, teamB *Team, date time.Time) *Match {
	return &Match{
		ID:    id,
		TeamA: teamA,
		TeamB: teamB,
		Date:  date,
	}
}