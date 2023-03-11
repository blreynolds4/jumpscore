package jumpscore

import (
	"sort"
)

type twoOfThree struct {
	getDistancePoints GetDistancePoints
}

func NewBestTwoOfThree(dp GetDistancePoints) CompetitionScoring {
	return &twoOfThree{
		getDistancePoints: dp,
	}
}

func (tt *twoOfThree) RoundCount() int {
	return 3
}

func (tt *twoOfThree) ScoreCompetition(c Competition) ScoredCompetition {
	// scores are provided in the order they happened
	// calculate the comp total by using the jumpers 2 highest scoring jumps
	scores := c.GetScores()

	// build a map of jumpers and their jumps
	jumpers := make(map[string][]ScoredJump)
	for _, jr := range scores {
		// make sure we have a list of scored jumps for each jumper
		scoredJumps, found := jumpers[jr.GetJumper().GetName()]
		if !found {
			scoredJumps = make([]ScoredJump, 0, c.GetRoundCount())
		}

		scoredJump := tt.scoreJump(jr)
		jumpers[jr.GetJumper().GetName()] = append(scoredJumps, scoredJump)
	}

	// get a total score for each jumper
	jumperResults := make([]JumperResult, 0, len(jumpers))
	for _, scores := range jumpers {
		jumper := scores[0].GetJump().GetJumper()
		total := float32(0)

		// sort each jumpers jumps by jump total
		rankedJumps := make([]ScoredJump, len(scores))
		// copy the jumps for sorting to preserve the order of scoring
		copy(rankedJumps, scores)
		sort.Sort(RankedJumpTotal(rankedJumps))
		// scores are limited to their best 2 jumps, first 2 in the list
		for _, jump := range rankedJumps[0:2] {
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

func (tt *twoOfThree) scoreJump(jr JumpResult) ScoredJump {
	// score the jump
	distPoints := tt.getDistancePoints(jr.GetJump().KPoint, jr.GetDistance())
	stylePoints := float32(0)
	for _, judge := range jr.GetJudgesScores() {
		stylePoints = stylePoints + judge.GetTotal()
	}

	return NewScoredJump(jr, distPoints, stylePoints)
}
