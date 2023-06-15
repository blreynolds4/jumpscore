package jumpscore

import "time"

// An Event represents a name with a list of competitions
// Like "The AOC Home Meet 2024"
// with the following competitions:
// U8 Girls on the K10
// U8 Boys on the K10
// Open Girls on the K10
// Open Boys on the K10
// etc
type Event struct {
	Name         string        `json:"name"`
	Date         time.Time     `json:"date"`
	Competitions []Competition `json:"competitions"`
}

func NewEvent(name string, day time.Time) Event {
	return Event{
		Name:         name,
		Date:         day,
		Competitions: make([]Competition, 0),
	}
}
