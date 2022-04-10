package objects

import "3d/internal/vector"

type Camera struct {
	Position  *vector.Vec3
	Direction *vector.Vec3
}

func NewCamera(pos, direct *vector.Vec3) *Camera {
	return &Camera{
		Position:  pos,
		Direction: direct,
	}
}
