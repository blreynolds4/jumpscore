package jumpscore

type SkiJumper interface {
	GetName() string
	GetAge() int
	GetAffliation() string
}

type CompetingJumper struct {
	Jumper SkiJumper `json:"jumper"`
	Bib    int       `json:"bib"`
}

type skiJumper struct {
	Name        string `json:"name"`
	Affiliation string `json:"affiliation"`
	Age         int    `json:"age"`
	Bib         int    `json:"bib"`
}

func NewJumper(name string, age int, affiliation string) SkiJumper {
	return &skiJumper{
		Name:        name,
		Affiliation: affiliation,
		Age:         age,
		Bib:         0,
	}
}

func (j *skiJumper) GetName() string {
	return j.Name
}

func (j *skiJumper) GetAge() int {
	return j.Age
}

func (j *skiJumper) GetAffliation() string {
	return j.Affiliation
}
