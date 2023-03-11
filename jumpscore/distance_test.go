package jumpscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleMultiplier(t *testing.T) {
	// using a multiplier
	distanceMultiplier := NewSimpleDistanceMultiplier(1.5)

	distance := float32(10)
	expectedPoints := float32(15)

	// hill size doesn't matter in this case but we provide it
	actual := distanceMultiplier(38, distance)

	assert.Equal(t, expectedPoints, actual)
}

func TestSixtyPlusBonusUnderK(t *testing.T) {

	sixtyPlusBonus := NewNewSixtyPlusBonus()
	distance := float32(8)
	expectedPoints := float32(56)

	// hill size doesn't matter in this case but we provide it
	actual := sixtyPlusBonus(10, distance)

	assert.Equal(t, expectedPoints, actual)
}

func TestSixtyPlusBonusAtK(t *testing.T) {

	sixtyPlusBonus := NewNewSixtyPlusBonus()
	distance := float32(10)
	expectedPoints := float32(60)

	// hill size doesn't matter in this case but we provide it
	actual := sixtyPlusBonus(10, distance)

	assert.Equal(t, expectedPoints, actual)
}

func TestSixtyPlusBonusOverK(t *testing.T) {

	sixtyPlusBonus := NewNewSixtyPlusBonus()
	distance := float32(15)
	expectedPoints := float32(70)

	// hill size doesn't matter in this case but we provide it
	actual := sixtyPlusBonus(10, distance)

	assert.Equal(t, expectedPoints, actual)
}
