package jumpscore

import (
	"fmt"
	"sort"
)

type CompetitionScoring interface {
	RoundCount() int
	ScoreCompetition(Competition) ScoredCompetition
}

type trialAndTwo struct {
	getDistancePoints GetDistancePoints
}

func NewTrialAndTwo(dp GetDistancePoints) CompetitionScoring {
	return &trialAndTwo{
		getDistancePoints: dp,
	}
}

func (tt *trialAndTwo) RoundCount() int {
	return 3
}

func (tt *trialAndTwo) ScoreCompetition(c Competition) ScoredCompetition {
	scores := c.GetScores()

	// build a map of jumpers and their 3 scored jumps
	jumpers := make(map[string][]ScoredJump)
	for _, jr := range scores {
		// m ake sure we have a list of scored jumps
		scoredJumps, found := jumpers[jr.GetJumper().GetName()]
		if !found {
			scoredJumps = make([]ScoredJump, 0)
		}

		// score the jump
		jump := jr.GetJump()
		distPoints := tt.getDistancePoints(jump.KPoint, jr.GetDistance())
		stylePoints := float32(0)
		for _, judge := range jr.GetJudgesScores() {
			stylePoints = stylePoints + judge.GetTotal()
		}

		scoredJump := NewScoredJump(jr, distPoints, stylePoints)
		jumpers[jr.GetJumper().GetName()] = append(scoredJumps, scoredJump)
	}

	// get a total score for each jumper
	jumperResults := make([]JumperResult, 0)
	for _, scores := range jumpers {
		// b, _ := json.MarshalIndent(scores, "", "  ")
		// fmt.Println("ALL", string(b))
		jumper := scores[0].GetJump().GetJumper()
		total := float32(0)
		// only include the last 2 jumps in the total
		for _, jump := range scores[1:] {
			fmt.Println("jumper", jump.GetJump().GetJumper().GetName(), "Dist", jump.GetJump().GetDistance(), "total", jump.GetTotalScore())
			total = total + jump.GetTotalScore()
		}

		jumperResult := NewJumperResult(jumper, total, scores)
		jumperResults = append(jumperResults, jumperResult)
	}

	// sort the result by total score and return
	sort.Sort(RankedByTotalPoints(jumperResults))

	// now create a scored competition with our results
	return NewScoredCompetition(c, jumperResults)
}
