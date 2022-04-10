package vector

import (
	"math"
)

//double sign(double a) { return (0 < a) - (a < 0); }
func Sign(i float64) float64 {
	if i > 0 {
		return 1
	}
	if i < 0 {
		return -1
	}
	return 0
}

//double step(double edge, double x) { return x > edge; }
func Step(edge, x float64) float64 {
	if x > edge {
		return 1.0
	}
	return 0.0
}

//float clamp(float value, float min, float max) { return fmax(fmin(value, max), min); }
func Clamp(value, min, max int) int {
	return int(math.Max(math.Min(float64(value), float64(max)), float64(min)))
}
