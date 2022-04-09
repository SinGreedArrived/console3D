package vector

import (
	"math"
)

func RotateX(a *Vec3, angle float64) *Vec3 {
	b := a
	//b.z = a.z * cos(angle) - a.y * sin(angle);
	b.Z = a.Z*math.Cos(angle) - a.Y*math.Sin(angle)
	//b.y = a.z * sin(angle) + a.y * cos(angle);
	b.Y = a.Z*math.Sin(angle) + a.Y*math.Cos(angle)
	return b
}

func RotateY(a *Vec3, angle float64) *Vec3 {
	b := a
	//b.x = a.x * cos(angle) - a.z * sin(angle);
	b.X = a.X*math.Cos(angle) - a.Z*math.Sin(angle)
	//b.z = a.x * sin(angle) + a.z * cos(angle);
	b.Z = a.X*math.Sin(angle) + a.Z*math.Cos(angle)
	return b
}

func RotateZ(a *Vec3, angle float64) *Vec3 {
	b := a
	//b.x = a.x * cos(angle) - a.y * sin(angle);
	b.X = a.X*math.Cos(angle) - a.Y*math.Sin(angle)
	//b.y = a.x * sin(angle) + a.y * cos(angle);
	b.Y = a.X*math.Sin(angle) + a.Y*math.Cos(angle)
	return b
}

//double sign(double a) { return (0 < a) - (a < 0); }
func Sign(i float64) float64 {
	var x, y float64
	if 0 < i {
		x = 1
	} else {
		x = 0
	}
	if i < 0 {
		y = 1
	} else {
		y = 0
	}
	return x - y
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

//float dot(vec3 const& a, vec3 const& b) { return a.x * b.x + a.y * b.y + a.z * b.z; }
func Dot3(a *Vec3, b interface{}) float64 {
	switch b.(type) {
	case *Vec3:
		return a.X*b.(*Vec3).X + a.Y*b.(*Vec3).Y + a.Z*b.(*Vec3).Z
	}
	return a.X*b.(float64) + a.Y*b.(float64) + a.Z*b.(float64)
}

//vec3 reflect(vec3 rd, vec3 n) { return rd - n * (2 * dot(n, rd)); }
func Reflect3(rd, n *Vec3) *Vec3 {
	return rd.Diff(n.Mult((2 * Dot3(n, rd))))
}

//vec3 step(vec3 const& edge, vec3 v) { return &Vec3(step(edge.x, v.x), step(edge.y, v.y), step(edge.z, v.z)); }
func Step3(edge interface{}, v *Vec3) *Vec3 {
	switch edge.(type) {
	case *Vec3:
		return &Vec3{
			X: Step(edge.(*Vec3).X, v.X),
			Y: Step(edge.(*Vec3).Y, v.Y),
			Z: Step(edge.(*Vec3).Z, v.Z),
		}
	}
	return &Vec3{
		X: Step(edge.(float64), v.X),
		Y: Step(edge.(float64), v.Y),
		Z: Step(edge.(float64), v.Z),
	}
}

//vec3 sign(vec3 const& v) { return &Vec3(sign(v.x), sign(v.y), sign(v.z)); }
func Sign3(v *Vec3) *Vec3 {
	return &Vec3{
		X: Sign(v.X),
		Y: Sign(v.Y),
		Z: Sign(v.Z),
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

func Sphere(ro, rd *Vec3, r float64) *Vec2 {
	b := Dot3(ro, rd)
	c := Dot3(ro, ro) - r*r
	h := b*b - c
	if h < 0.0 {
		return NewVec2(-1.0, -1.0)
	}
	h = math.Sqrt(h)
	return NewVec2(-b-h, -b+h)
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

func Box(ro, rd, boxSize *Vec3, outNormal *Vec3) *Vec2 {
	m := NewVec3(1).Div(rd)
	n := m.Mult(ro)
	k := m.Abs().Mult(boxSize)
	t1 := n.Minus().Diff(k)
	t2 := n.Minus().Sum(k)
	tN := math.Max(math.Max(t1.X, t1.Y), t1.Z)
	tF := math.Min(math.Min(t2.X, t2.Y), t2.Z)
	if tN > tF || tF < 0.0 {
		return NewVec2(-1.0, -1.0)
	}
	yzx := NewVec3(t1.Y, t1.Z, t1.X)
	zxy := NewVec3(t1.Z, t1.X, t1.Y)
	outNormal = Sign3(rd).Mult(Step3(yzx, t1)).Mult(Step3(zxy, t1)).Minus()
	return NewVec2(tN, tF)
}

/*
float plane(vec3 ro, vec3 rd, vec3 p, float w) {
    return -(dot(ro, p) + w) / dot(rd, p);
}
*/
func Plane(ro, rd, p *Vec3, w float64) float64 {
	return -(Dot3(ro, p) + w) / Dot3(rd, p)
}
