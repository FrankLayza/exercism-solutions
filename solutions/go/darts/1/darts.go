package darts

import "math"

func Score(x, y float64) int {
	var points int
	distance := math.Sqrt((x * x + y * y))
	if distance <= 1{
		points += 10
	}else if distance <= 5{
		points += 5
	}else if distance <= 10{
		points += 1
	}else{
		points = 0
	}
	return points
}
