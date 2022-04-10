package objects

import (
	"3d/internal/vector"
	"math"
)

type Sphere struct {
	Position *vector.Vec3
	Radius   float64
}

func NewSphere(pos *vector.Vec3, radius float64) *Sphere {
	return &Sphere{
		Position: pos,
		Radius:   radius,
	}
}

/*
vec2 sphere(vec3 ro, vec3 rd, float r) {
    float b = dot(ro, rd);
    float c = dot(ro, ro) - r * r;
    float h = b * b - c;
    if (h < 0.0) return &Vec2(-1.0);
    h = sqrt(h);
    return &Vec2(-b - h, -b + h);
}
*/
func (s *Sphere) Intersection(camera *Camera) *vector.Vec2 {
	diff := camera.Position.Diff(s.Position)
	b := diff.Dot(camera.Direction)
	c := diff.Dot(diff) - s.Radius*s.Radius
	h := b*b - c
	if h < 0.0 {
		return vector.NewVec2(-1.0, -1.0)
	}
	h = math.Sqrt(h)
	return vector.NewVec2(-b-h, -b+h)
}

func (s *Sphere) SetRadius(i float64) {
	s.Radius = i
}
func (s *Sphere) GetRadius() float64 {
	return s.Radius
}
