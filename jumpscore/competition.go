package jumpscore

import "fmt"

type Competition interface {
	GetRoundCount() int
	AddJumper(SkiJumper)
	JumperOrder() ([]CompetingJumper, error)
	AddScore(int, CompetingJumper, float32, []StyleScore)
	GetScores() []JumpResult
}

type competition struct {
	Name      string      `json:"name"`
	Jump      SkiJump     `json:"skiJump"`
	Bibs      BibSet      `json:"bibs"`
	Rounds    int         `json:"rounds"`
	Jumpers   []SkiJumper `json:"jumpers"`
	jumpOrder []CompetingJumper
	Scores    []JumpResult `json:"rawScores"`
}

func NewCompetition(name string, jump SkiJump, bibs BibSet, rounds int) Competition {
	return &competition{
		Name:      name,
		Jump:      jump,
		Bibs:      bibs,
		Rounds:    rounds,
		Jumpers:   make([]SkiJumper, 0, 5),
		jumpOrder: nil,
		Scores:    nil,
	}
}

func (c *competition) GetRoundCount() int {
	return c.Rounds
}

func (c *competition) AddJumper(jumper SkiJumper) {
	// find and update if the jumper exists
	found := -1
	for i := 0; i < len(c.Jumpers); i++ {
		if jumper.GetName() == c.Jumpers[i].GetName() {
			found = i
		}
	}

	if found > -1 {
		c.Jumpers[found] = jumper
	} else {
		c.Jumpers = append(c.Jumpers, jumper)
	}
}

func (c *competition) JumperOrder() ([]CompetingJumper, error) {
	// order is in the order of the additions
	// assign bibs to jumpers from the bib set

	// can only be set once
	if c.jumpOrder != nil {
		return c.jumpOrder, nil
	}

	bibs := c.Bibs.AvailableBibs()
	if len(bibs) < len(c.Jumpers) {
		return nil, fmt.Errorf("not enough bibs for jumpers:  %d bibs for %d jumpers", len(bibs), len(c.Jumpers))
	}

	c.jumpOrder = make([]CompetingJumper, len(c.Jumpers))
	for i, jumper := range c.Jumpers {
		c.jumpOrder[i].Jumper = jumper
		c.jumpOrder[i].Bib = bibs[i]
		c.Bibs.Assign(bibs[i])
	}

	return c.jumpOrder, nil
}

func (c *competition) AddScore(round int, jumper CompetingJumper, dist float32, style []StyleScore) {
	if c.Scores == nil {
		// make enough results room for jump result for each round
		c.Scores = make([]JumpResult, 0, len(c.jumpOrder)*c.Rounds)
	}

	newResult := newJumpResult(round, c.Jump, jumper.Bib, jumper.Jumper, dist, style)

	// prevent and/or overwrite duplicate results (1 per jumper per round)
	for i, score := range c.Scores {
		if round == score.GetRound() && jumper.Bib == score.GetBib() {
			// update the existing score and return
			c.Scores[i] = newResult
			return
		}
	}

	// result was not found add it on
	c.Scores = append(c.Scores, newResult)
}

func (c *competition) GetScores() []JumpResult {
	result := make([]JumpResult, 0)
	// make a dump copy of the results
	return append(result, c.Scores...)
}
