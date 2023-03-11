package jumpscore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoringTrialAndTwo2023NHIAA(t *testing.T) {
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

	trialAndTwo := NewTrialAndTwo(NewSimpleDistanceMultiplier(float32(1.54)))
	scoredComp := trialAndTwo.ScoreCompetition(comp)
	jumperResults := scoredComp.GetJumperResults()

	// ORDER OF JUMPS (Round 1, 2, 3) not being maintained

	for place, jr := range jumperResults {
		fmt.Printf("%d\t%s\t%f\n", place+1, jr.GetJumper().GetName(), jr.GetTotalScore())
	}

	// assert the expected comp results
	assert.Equal(t, jumperResults[0].GetJumper().GetName(), schuyler.GetName())
	assert.InDelta(t, float32(220.98), jumperResults[0].GetTotalScore(), .01)

	assert.Equal(t, jumperResults[1].GetJumper().GetName(), angelo.GetName())
	assert.InDelta(t, float32(209.74), jumperResults[1].GetTotalScore(), .01)

	assert.Equal(t, jumperResults[2].GetJumper().GetName(), mychal.GetName())
	assert.InDelta(t, float32(204.62), jumperResults[2].GetTotalScore(), .01)
}
