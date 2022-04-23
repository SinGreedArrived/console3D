package objects

import "3d/internal/vector"

type Screen interface {
	UV() *vector.Vec2
	SetPixel(float64)
	NextCoord() bool
	Render()
}
