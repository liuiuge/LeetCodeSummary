package code

import (
	// "math/rand"
	"time"
)

// https://leetcode.com/problems/generate-random-point-in-a-circle/

type CSolution struct {
	x_center float64
	y_center float64
	radius   float64
}

func CirConstructor(radius float64, x_center float64, y_center float64) CSolution {
	return CSolution{
		x_center: x_center,
		y_center: y_center,
		radius:   radius,
	}
}

func (this *CSolution) RandPoint() []float64 {
	return []float64{this.x_center + this.radius/float64(time.Now().Unix()%10), this.y_center}
}
