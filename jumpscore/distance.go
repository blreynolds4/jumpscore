package jumpscore

type GetDistancePoints func(kPoint float32, distance float32) float32

func NewSimpleDistanceMultiplier(multiplier float32) GetDistancePoints {
	return func(hillSize float32, distance float32) float32 {
		return distance * multiplier
	}
}

func NewNewSixtyPlusBonus() GetDistancePoints {
	return func(kPoint float32, distance float32) float32 {
		bonusDist := distance - kPoint
		// 60 point bonus + 2 points per meter over K, -2 points per meter under K
		return 60 + 2*bonusDist
	}
}
