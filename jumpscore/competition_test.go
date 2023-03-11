package jumpscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalCompetition(t *testing.T) {
	// build test comp
	sj := NewSkiJump(t.Name(), 90)
	bibs := NewBibSet(1, 3, []int{})
	comp := NewCompetition(t.Name(), sj, bibs, 3)
	assert.Equal(t, 3, comp.GetRoundCount())

	// create and add 2 jumpers
	jumper1 := NewJumper(t.Name()+"1", 15, "abc")
	jumper2 := NewJumper(t.Name()+"2", 25, "def")
	comp.AddJumper(jumper1)
	comp.AddJumper(jumper2)

	order, err := comp.JumperOrder()
	assert.NoError(t, err)

	assert.Equal(t, 2, len(order))
	// verify bib 3 is still available
	assert.Equal(t, 1, len(bibs.AvailableBibs()))
	assert.Equal(t, 3, bibs.AvailableBibs()[0])

	// verify we don't assign more bibs with second call to jump order
	dupOrder, err := comp.JumperOrder()
	assert.NoError(t, err)

	assert.Equal(t, 2, len(dupOrder))
	// verify bib 3 is still available
	assert.Equal(t, 1, len(bibs.AvailableBibs()))
	assert.Equal(t, 3, bibs.AvailableBibs()[0])

	// verify jumpers got bibs 1 and 2
	assert.Equal(t, jumper1.GetName(), order[0].Jumper.GetName())
	assert.Equal(t, 1, order[0].Bib)

	assert.Equal(t, jumper2.GetName(), order[1].Jumper.GetName())
	assert.Equal(t, 2, order[1].Bib)

	// add some scores
	comp.AddScore(1, order[0], 0.5, []StyleScore{})
	comp.AddScore(1, order[0], 90.5, []StyleScore{})
	comp.AddScore(1, order[1], 50, []StyleScore{})

	scores := comp.GetScores()
	assert.Equal(t, 2, len(scores)) // no dups, keep latest
	assert.Equal(t, order[0].Bib, scores[0].GetBib())
	assert.Equal(t, float32(90.5), scores[0].GetDistance())

	assert.Equal(t, order[1].Bib, scores[1].GetBib())
	assert.Equal(t, float32(50), scores[1].GetDistance())
}

func TestJumperOrderDupJumper(t *testing.T) {
	// build test comp
	sj := NewSkiJump(t.Name(), 90)
	bibs := NewBibSet(1, 3, []int{})
	comp := NewCompetition(t.Name(), sj, bibs, 2)

	// create and add 2 jumpers
	jumper1 := NewJumper(t.Name()+"1", 15, "abc")
	comp.AddJumper(jumper1)
	jumper1 = NewJumper(t.Name()+"1", 15, "abc")
	comp.AddJumper(jumper1)

	order, err := comp.JumperOrder()
	assert.NoError(t, err)

	assert.Equal(t, 1, len(order))
	// verify bib 2 bibs still available
	assert.Equal(t, 2, len(bibs.AvailableBibs()))

	// verify jumper got bibs 1
	assert.Equal(t, jumper1.GetName(), order[0].Jumper.GetName())
	assert.Equal(t, jumper1.GetAge(), order[0].Jumper.GetAge())
	assert.Equal(t, 1, order[0].Bib)
}

func TestJumperOrderNotEnoughBibs(t *testing.T) {
	// build test comp
	sj := NewSkiJump(t.Name(), 90)
	bibs := NewBibSet(1, 1, []int{})
	comp := NewCompetition(t.Name(), sj, bibs, 1)

	// create and add 2 jumpers
	jumper1 := NewJumper(t.Name()+"1", 15, "abc")
	jumper2 := NewJumper(t.Name()+"2", 25, "def")
	comp.AddJumper(jumper1)
	comp.AddJumper(jumper2)

	order, err := comp.JumperOrder()
	assert.Nil(t, order)
	assert.Error(t, err)
	assert.Equal(t, "not enough bibs for jumpers:  1 bibs for 2 jumpers",
		err.Error())
}
