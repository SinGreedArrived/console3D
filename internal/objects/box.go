package objects

import (
	"3d/internal/vector"
	"math"
)

type Box struct {
	Position *vector.Vec3
	Size     *vector.Vec3
}

func NewBox(pos, size *vector.Vec3) *Box {
	return &Box{
		Position: pos,
		Size:     size,
	}
}

/*
vec2 box(vec3 ro, vec3 rd, vec3 boxSize, vec3& outNormal) {
    vec3 m = vec3(1.0) / rd;
    vec3 n = m * ro;
    vec3 k = abs(m) * boxSize;
    vec3 t1 = -n - k;
    vec3 t2 = -n + k;
    float tN = fmax(fmax(t1.x, t1.y), t1.z);
    float tF = fmin(fmin(t2.x, t2.y), t2.z);
    if (tN > tF || tF < 0.0) {
        return &Vec2(-1.0);
    };
    std::cout << n.x << " " << n.y << " " << n.z << "\n";
    vec3 yzx = vec3(t1.y, t1.z, t1.x);
    vec3 zxy = vec3(t1.z, t1.x, t1.y);
    outNormal = -sign(rd) * step(yzx, t1) * step(zxy, t1);
    return &Vec2(tN, tF);
}
*/

func (b *Box) Intersection(camera *Camera) (*vector.Vec2, *vector.Vec3) {
	diff := camera.Position.Diff(b.Position)
	m := vector.NewVec3(1).Div(camera.Direction)
	n := m.Mult(diff)
	k := m.Abs().Mult(b.Size)
	t1 := n.Minus().Diff(k)
	t2 := n.Minus().Sum(k)
	tN := math.Max(math.Max(t1.X, t1.Y), t1.Z)
	tF := math.Min(math.Min(t2.X, t2.Y), t2.Z)
	if tN > tF || tF < 0.0 {
		return vector.NewVec2(-1.0, -1.0), nil
	}
	yzx := vector.NewVec3(t1.Y, t1.Z, t1.X)
	zxy := vector.NewVec3(t1.Z, t1.X, t1.Y)
	outNormal := camera.Direction.Sign().Mult(t1.Step(yzx)).Mult(t1.Step(zxy)).Minus()
	return vector.NewVec2(tN, tF), outNormal.Norm()
}
