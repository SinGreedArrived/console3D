package objects

import "3d/internal/vector"

type Screen interface {
	UV(*vector.Vec2) *vector.Vec2
	SetPixel(float64, *vector.Vec2)
	NextCoord() bool
	Render()
	GetSize() (uint64, uint64)
	GetCoord() *vector.Vec2
}
