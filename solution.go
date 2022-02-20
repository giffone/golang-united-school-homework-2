package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sides int

const (
	SidesTriangle sides = 3
	SidesSquare   sides = 4
	SidesCircle   sides = 0
)

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	if sidesNum == 3 {
		sRoot := math.Sqrt(3)
		root := sideLen * sideLen
		return (sRoot / 4) * root
	}
	if sidesNum == 0 {
		pi := math.Pi * 4
		root := sideLen * sideLen
		return root / pi
	}
	if sidesNum == 4 {
		return sideLen * sideLen
	}
	return 0
}
