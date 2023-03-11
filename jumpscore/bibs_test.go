package jumpscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalBibSetNoMissing(t *testing.T) {
	start := 1
	end := 10
	bibSet := NewBibSet(start, end, []int{})

	available := bibSet.AvailableBibs()
	assert.Equal(t, 10, len(available))

	// bibs should be in numeric order
	for i := start; i <= end; i++ {
		assert.True(t, available[i-1] == i)

		// assign the bibs
		bibSet.Assign(available[i-1])
	}

	available = bibSet.AvailableBibs()
	assert.Equal(t, 0, len(available))

	bibSet.ClearAssignments()
	available = bibSet.AvailableBibs()
	assert.Equal(t, 10, len(available))
}

func TestNormalBibSetMissing(t *testing.T) {
	start := 1
	end := 10
	bibSet := NewBibSet(start, end, []int{5})

	available := bibSet.AvailableBibs()
	assert.Equal(t, 9, len(available))
}
