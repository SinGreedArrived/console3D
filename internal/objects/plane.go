package objects

import "3d/internal/vector"

type Plane struct {
	Position *vector.Vec3
}

func NewPlane(pos *vector.Vec3) *Plane {
	return &Plane{
		Position: pos,
	}
}

/*
float plane(vec3 ro, vec3 rd, vec3 p, float w) {
    return -(dot(ro, p) + w) / dot(rd, p);
}
*/
func (p *Plane) Intersection(camera *Camera) *vector.Vec2 {
	return vector.NewVec2(-(camera.Position.Dot(p.Position) + 1) / camera.Direction.Dot(p.Position))
}
