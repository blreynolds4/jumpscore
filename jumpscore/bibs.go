package jumpscore

import "sort"

type BibSet interface {
	AvailableBibs() []int
	Assign(bib int)
	ClearAssignments()
}

type bibSet struct {
	bibs map[int]bool
}

func (bs *bibSet) AvailableBibs() []int {
	result := make(sort.IntSlice, 0)

	for bib, assigned := range bs.bibs {
		if !assigned {
			result = append(result, bib)
		}
	}

	result.Sort()

	return result
}

func (bs *bibSet) Assign(bib int) {
	bs.bibs[bib] = true
}

func (bs *bibSet) ClearAssignments() {
	for bib := range bs.bibs {
		bs.bibs[bib] = false
	}
}

func NewBibSet(rangeStart int, rangeEnd int, missingBibs []int) BibSet {

	result := &bibSet{
		bibs: make(map[int]bool),
	}

	for i := rangeStart; i <= rangeEnd; i++ {
		// bib is available in the set
		result.bibs[i] = false
	}

	// take out the missing bibs
	for _, missingBib := range missingBibs {
		// remove bibs from the set
		// (not just assigned)
		delete(result.bibs, missingBib)
	}

	return result
}
