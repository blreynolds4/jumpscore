package main

import (
	"fmt"
	"jumpscore/jumpscore"
)

func main() {
	// set up a hill and competition (class) on the hill
	k38 := jumpscore.NewSkiJump("K38", 38)

	// create a set of bibs given a range and list of missing numbers
	// should share a bib set across all competitions
	// bibs can be marked assigned so each comp gets part of the range
	bibSet := jumpscore.NewBibSet(1, 20, []int{2, 6, 7, 8})

	// create a scoring method.  We need a distance points calculation to create.
	// The scoring method will
	// generate a round count.  It will accept a list of results
	// from the competition and produce final results
	// initial scoring methods are trial and 2 or best 2 out of 3
	// (also do longest standing)
	trialAnd2 := jumpscore.NewTrialAndTwo(jumpscore.NewSimpleDistanceMultiplier(float32(1.54)))

	// Create a competition with name, jump and bibs
	// Competition is for each class
	k38hs := jumpscore.NewCompetition("HS Class", k38, bibSet, 3)
	// k38u14 := jumpscore.NewCompetition("U14", k38, bibSet)
	// k38Seniors := jumpscore.NewCompetition("Seniors", k38, bibSet)

	// add athletes to the competition
	k38hs.AddJumper(jumpscore.NewJumper("Mychal", 15, "AOC"))
	k38hs.AddJumper(jumpscore.NewJumper("Angelo", 15, "AOC"))
	k38hs.AddJumper(jumpscore.NewJumper("Schyler", 15, "AOC"))

	// get a comp order with bib assigned
	// the result is a list of jumpers
	jumperOrder, _ := k38hs.JumperOrder()

	// run the comp based on round count
	for i := 1; i <= k38hs.GetRoundCount(); i++ {
		// run through the jump order for each round
		// add scores for each jumper to the comp
		for _, bibJumper := range jumperOrder {
			styleScores := []jumpscore.StyleScore{
				jumpscore.NewStyleScore("John", 13),
				jumpscore.NewStyleScore("Mark", 13.5),
				jumpscore.NewStyleScore("Walter", 14),
			}
			k38hs.AddScore(i, bibJumper, 34, styleScores)
		}
	}

	// score the competition
	competitionResults := trialAnd2.ScoreCompetition(k38hs)

	// print results
	for place, result := range competitionResults.GetJumperResults() {
		fmt.Printf("%d\t%s, %s\n", place+1, result.GetJumper().GetName(), result.GetJumper().GetAffliation())
	}
}
