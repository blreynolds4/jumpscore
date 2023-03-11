package jumpscore

type SkiJump struct {
	JumpName string  `json:"name"`
	KPoint   float32 `json:"kPoint"`
}

func NewSkiJump(name string, k float32) SkiJump {
	return SkiJump{
		JumpName: name,
		KPoint:   k,
	}
}
