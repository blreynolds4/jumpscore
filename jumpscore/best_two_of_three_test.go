package jumpscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoringBest2of3_2023NHIAA(t *testing.T) {
	// build test comp
	sj := NewSkiJump(t.Name(), 39)
	bibs := NewBibSet(1, 3, []int{})
	comp := NewCompetition(t.Name(), sj, bibs, 3)

	// create and add 3 jumpers
	mychal := NewJumper("Mychal", 15, "AOC")
	angelo := NewJumper("Angelo", 15, "LOC")
	schuyler := NewJumper("Schuyler", 15, "NYSEF")
	comp.AddJumper(mychal)
	comp.AddJumper(angelo)
	comp.AddJumper(schuyler)

	order, err := comp.JumperOrder()
	assert.NoError(t, err)

	// add some scores
	//  mychal
	comp.AddScore(1, order[0], 39, []StyleScore{NewStyleScore("all", 44.5)})
	comp.AddScore(2, order[0], 39, []StyleScore{NewStyleScore("all", 44.5)})
	comp.AddScore(3, order[0], 39, []StyleScore{NewStyleScore("all", 40)})

	// angelo
	comp.AddScore(1, order[1], 43, []StyleScore{NewStyleScore("all", 41)})
	comp.AddScore(2, order[1], 41.5, []StyleScore{NewStyleScore("all", 42.5)})
	comp.AddScore(3, order[1], 39.5, []StyleScore{NewStyleScore("all", 42.5)})

	// schuyler
	comp.AddScore(1, order[2], 43.5, []StyleScore{NewStyleScore("all", 45)})
	comp.AddScore(2, order[2], 43.5, []StyleScore{NewStyleScore("all", 45)})
	comp.AddScore(3, order[2], 43.5, []StyleScore{NewStyleScore("all", 42)})

	bestTwo := NewBestTwoOfThree(NewSimpleDistanceMultiplier(float32(1.54)))
	scoredComp := bestTwo.ScoreCompetition(comp)
	jumperResults := scoredComp.GetJumperResults()

	// assert the expected comp results
	assert.Equal(t, jumperResults[0].GetJumper().GetName(), schuyler.GetName())
	assert.InDelta(t, float32(223.98), jumperResults[0].GetTotalScore(), .001)

	assert.Equal(t, jumperResults[1].GetJumper().GetName(), angelo.GetName())
	assert.InDelta(t, float32(213.63), jumperResults[1].GetTotalScore(), .001)

	assert.Equal(t, jumperResults[2].GetJumper().GetName(), mychal.GetName())
	assert.InDelta(t, float32(209.12), jumperResults[2].GetTotalScore(), .001)

	// assert we have enough info to do full competition report
	assert.Equal(t, 3, len(jumperResults))
	// each jumper should have 3 results
	for _, jumper := range jumperResults {
		assert.Equal(t, 3, len(jumper.GetScoredJumps()))
	}
}
