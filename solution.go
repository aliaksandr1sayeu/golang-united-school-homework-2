package square

import "math"

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

type anotherInt int

func CalcSquare(sideLen float64, sidesNum anotherInt) float64 {
	switch sidesNum {
	case SidesTriangle:
		return triangleArea(sideLen)
	case SidesSquare:
		return squareArea(sideLen)
	case SidesCircle:
		return circleArea(sideLen)
	default:
		return 0
	}
}

func triangleArea(sideLen float64) float64 {
	return math.Sqrt(3) * math.Pow(sideLen, 2) / 4
}
func squareArea(sideLen float64) float64 {
	return math.Pow(sideLen, 2)
}
func circleArea(sideLen float64) float64 {
	return math.Pi * math.Pow(sideLen, 2) / 4
}
