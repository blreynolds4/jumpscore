package jumpscore

type StyleScore interface {
	GetJudge() string
	GetTotal() float32
}

type styleScore struct {
	Judge string  `json:"judge"`
	Total float32 `json:"totalScore"`
}

func NewStyleScore(judge string, total float32) StyleScore {
	return &styleScore{
		Judge: judge,
		Total: total,
	}
}

func (ss *styleScore) GetJudge() string {
	return ss.Judge
}

func (ss *styleScore) GetTotal() float32 {
	return ss.Total
}

type JumpResult interface {
	GetRound() int
	GetJump() SkiJump
	GetBib() int
	GetJumper() SkiJumper
	GetDistance() float32
	GetJudgesScores() []StyleScore
}

type jumpResult struct {
	Round        int          `json:"round"`
	Bib          int          `json:"bib"`
	Jumper       SkiJumper    `json:"jumper"`
	Distance     float32      `json:"distance"`
	JudgesScores []StyleScore `json:"scores"`
	Jump         SkiJump      `json:"jump"`
}

func newJumpResult(round int, jump SkiJump, bib int, jumper SkiJumper, distance float32, judges []StyleScore) JumpResult {
	return &jumpResult{
		Round:        round,
		Jumper:       jumper,
		Bib:          bib,
		Distance:     distance,
		JudgesScores: judges,
		Jump:         jump,
	}
}

func (jr *jumpResult) GetRound() int {
	return jr.Round
}

func (jr *jumpResult) GetBib() int {
	return jr.Bib
}

func (jr *jumpResult) GetJumper() SkiJumper {
	return jr.Jumper
}

func (jr *jumpResult) GetJump() SkiJump {
	return jr.Jump
}

func (jr *jumpResult) GetDistance() float32 {
	return jr.Distance
}

func (jr *jumpResult) GetJudgesScores() []StyleScore {
	return jr.JudgesScores
}
