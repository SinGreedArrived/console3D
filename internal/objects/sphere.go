package objects

import "3d/internal/vector"

type Sphere struct {
	vector.Vec3
}

func NewSphere(pos *vector.Vec3) *Sphere {
	return &Sphere{
		Vec3: *pos,
	}
}
