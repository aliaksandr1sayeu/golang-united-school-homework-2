package square

import "math"

const (
	sidesTriangle = 3
	sidesSquare   = 4
	sidesCircle   = 0
)

type anotherInt int

func CalcSquare(sideLen float64, sidesNum anotherInt) float64 {
	switch sidesNum {
	case sidesTriangle:
		return triangleArea(sideLen)
	case sidesSquare:
		return squareArea(sideLen)
	case sidesCircle:
		return circleArea(sideLen)
	default:
		return 0
	}
}

func triangleArea(sideLen float64) float64 {
	return math.Sqrt(0.75) * math.Pow(sideLen, 2)
}
func squareArea(sideLen float64) float64 {
	return math.Pow(sideLen, 2)
}
func circleArea(sideLen float64) float64 {
	return math.Pi * math.Pow(sideLen, 2)
}
