package jumpscore

type ScoredJump interface {
	GetJump() JumpResult
	GetDistanceScore() float32
	GetStyleTotal() float32
	GetTotalScore() float32
}

type JumperScores interface {
	GetJumper() SkiJumper
	GetScoredJumps() []ScoredJump
}

type JumperResult interface {
	GetTotalScore() float32
	GetJumper() SkiJumper
	GetScoredJumps() []ScoredJump
}

type ScoredCompetition interface {
	GetCompetition() Competition
	GetJumperResults() []JumperResult
}

type scoredJump struct {
	JumpScore     JumpResult
	DistanceScore float32
	StyleScore    float32
}

func (sj *scoredJump) GetJump() JumpResult {
	return sj.JumpScore
}

func (sj *scoredJump) GetDistanceScore() float32 {
	return sj.DistanceScore
}

func (sj *scoredJump) GetStyleTotal() float32 {
	return sj.StyleScore
}

func (sj *scoredJump) GetTotalScore() float32 {
	return sj.DistanceScore + sj.StyleScore
}

func NewScoredJump(jr JumpResult, distPoints float32, stylePoints float32) ScoredJump {
	return &scoredJump{
		JumpScore:     jr,
		DistanceScore: distPoints,
		StyleScore:    stylePoints,
	}
}

type jumperResult struct {
	Jumper     SkiJumper    `json:"jumper"`
	TotalScore float32      `json:"totalScore"`
	Jumps      []ScoredJump `json:"jumps"`
}

func (jr *jumperResult) GetTotalScore() float32 {
	return jr.TotalScore
}

func (jr *jumperResult) GetJumper() SkiJumper {
	return jr.Jumper
}

func (jr *jumperResult) GetScoredJumps() []ScoredJump {
	return jr.Jumps
}

func NewJumperResult(jumper SkiJumper, total float32, jumps []ScoredJump) JumperResult {
	return &jumperResult{
		Jumper:     jumper,
		TotalScore: total,
		Jumps:      jumps,
	}
}

// ByJumpTotal implements sort.Interface for []ScoredJump based on
// TotalPoints.
type RankedJumpTotal []ScoredJump

func (rjt RankedJumpTotal) Len() int {
	return len(rjt)
}

func (rjt RankedJumpTotal) Swap(i, j int) {
	rjt[i], rjt[j] = rjt[j], rjt[i]
}

func (rjt RankedJumpTotal) Less(i, j int) bool {
	return rjt[i].GetTotalScore() > rjt[j].GetTotalScore()
}

// ByTotalPoints implements sort.Interface for []JumperResult based on
// TotalPoints.
type RankedByTotalPoints []JumperResult

func (jrs RankedByTotalPoints) Len() int {
	return len(jrs)
}

func (jrs RankedByTotalPoints) Swap(i, j int) {
	jrs[i], jrs[j] = jrs[j], jrs[i]
}

func (jrs RankedByTotalPoints) Less(i, j int) bool {
	return jrs[i].GetTotalScore() > jrs[j].GetTotalScore()
}

type scoredCompetition struct {
	Comp            Competition
	DetailedResults []JumperResult
	Results         []JumperResult
}

func (sc *scoredCompetition) GetCompetition() Competition {
	return sc.Comp
}

/*
Get the Jumpers with scores for all their jumps.
*/
func (sc *scoredCompetition) GetDetailedJumperResults() []JumperResult {
	return sc.DetailedResults
}

/*
Get the Jumpers with the jumps that made up their scores.  This may be
a different (smaller) set of jumps than the detailed results.
*/
func (sc *scoredCompetition) GetJumperResults() []JumperResult {
	return sc.Results
}

func NewScoredCompetition(c Competition, results []JumperResult) ScoredCompetition {
	return &scoredCompetition{
		Comp:    c,
		Results: results,
	}
}
