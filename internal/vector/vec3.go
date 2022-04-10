package vector

import (
	. "math"
	//"github.com/go-gl/gl/v2.1/gl"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func NewVec3(i ...interface{}) *Vec3 {
	return newVec3(i)
}
func newVec3(i []interface{}) *Vec3 {
	v := &Vec3{}
	if len(i) == 1 {
		switch i[0].(type) {
		case int:
			return &Vec3{
				X: float64(i[0].(int)),
				Y: float64(i[0].(int)),
				Z: float64(i[0].(int)),
			}
		case float64:
			return &Vec3{
				X: i[0].(float64),
				Y: i[0].(float64),
				Z: i[0].(float64),
			}
		}
	}
	for n := range i {
		switch i[n].(type) {
		case int:
			switch n {
			case 0:
				v.X = float64(i[n].(int))
			case 1:
				v.Y = float64(i[n].(int))
			case 2:
				v.Z = float64(i[n].(int))
			}
		case float64:
			switch n {
			case 0:
				v.X = i[n].(float64)
			case 1:
				v.Y = i[n].(float64)
			case 2:
				v.Z = i[n].(float64)
			}
		}
	}
	return v
}

func (v *Vec3) Sum(i interface{}) *Vec3 {
	switch i.(type) {
	case *Vec3:
		vector3 := i.(*Vec3)
		return &Vec3{
			X: v.X + vector3.X,
			Y: v.Y + vector3.Y,
			Z: v.Z + vector3.Z,
		}
	case float64:
		f := i.(float64)
		return &Vec3{
			X: v.X + f,
			Y: v.Y + f,
			Z: v.Z + f,
		}
	}
	return nil
}

func (v *Vec3) Diff(i interface{}) *Vec3 {
	switch i.(type) {
	case *Vec3:
		vector3 := i.(*Vec3)
		return &Vec3{
			X: v.X - vector3.X,
			Y: v.Y - vector3.Y,
			Z: v.Z - vector3.Z,
		}
	case float64:
		f := i.(float64)
		return &Vec3{
			X: v.X - f,
			Y: v.Y - f,
			Z: v.Z - f,
		}
	}
	return nil
}

func (v *Vec3) Minus() *Vec3 {
	return &Vec3{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func (v *Vec3) Mult(i interface{}) *Vec3 {
	switch i.(type) {
	case *Vec3:
		vector3 := i.(*Vec3)
		return &Vec3{
			X: v.X * vector3.X,
			Y: v.Y * vector3.Y,
			Z: v.Z * vector3.Z,
		}
	case float64:
		f := i.(float64)
		return &Vec3{
			X: v.X * f,
			Y: v.Y * f,
			Z: v.Z * f,
		}
	}
	return nil
}

func (v *Vec3) Div(i interface{}) *Vec3 {
	switch i.(type) {
	case *Vec3:
		vector3 := i.(*Vec3)
		return &Vec3{
			X: v.X / vector3.X,
			Y: v.Y / vector3.Y,
			Z: v.Z / vector3.Z,
		}
	case float64:
		f := i.(float64)
		return &Vec3{
			X: v.X / f,
			Y: v.Y / f,
			Z: v.Z / f,
		}
	}
	return nil
}

func (v *Vec3) Step(i interface{}) *Vec3 {
	switch i.(type) {
	case *Vec3:
		vector3 := i.(*Vec3)
		return &Vec3{
			X: Step(vector3.X, v.X),
			Y: Step(vector3.Y, v.Y),
			Z: Step(vector3.Z, v.Z),
		}
	case float64:
		f := i.(float64)
		return &Vec3{
			X: Step(f, v.X),
			Y: Step(f, v.Y),
			Z: Step(f, v.Z),
		}
	}
	return nil
}

func (v *Vec3) Abs() *Vec3 {
	return &Vec3{
		X: Abs(v.X),
		Y: Abs(v.Y),
		Z: Abs(v.Z),
	}
}

func (v *Vec3) Sign() *Vec3 {
	return &Vec3{
		X: Sign(v.X),
		Y: Sign(v.Y),
		Z: Sign(v.Z),
	}
}

func (v *Vec3) Reflect(n *Vec3) *Vec3 {
	return v.Diff(n.Mult(n.Dot(v) * 2))
}

func (v *Vec3) Len() float64 {
	return Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vec3) Norm() *Vec3 {
	return v.Div(v.Len())
}

func (v *Vec3) Dot(i interface{}) float64 {
	switch i.(type) {
	case *Vec3:
		vector3 := i.(*Vec3)
		return v.X*vector3.X + v.Y*vector3.Y + v.Z*vector3.Z
	case float64:
		f := i.(float64)
		return v.X*f + v.Y*f + v.Z*f
	}
	return 0.0
}

func (v *Vec3) RotateX(angle float64) {
	tmpV := *v
	//b.z = a.z * cos(angle) - a.y * sin(angle);
	v.Z = tmpV.Z*Cos(angle) - tmpV.Y*Sin(angle)
	//b.y = a.z * sin(angle) + a.y * cos(angle);
	v.Y = tmpV.Z*Sin(angle) + tmpV.Y*Cos(angle)
}

func (v *Vec3) RotateY(angle float64) {
	tmpV := *v
	//b.x = a.x * cos(angle) - a.z * sin(angle);
	v.X = tmpV.X*Cos(angle) - tmpV.Z*Sin(angle)
	//b.z = a.x * sin(angle) + a.z * cos(angle);
	v.Z = tmpV.X*Sin(angle) + tmpV.Z*Cos(angle)
}

func (v *Vec3) RotateZ(angle float64) {
	tmpV := *v
	//b.x = a.x * cos(angle) - a.y * sin(angle);
	v.X = tmpV.X*Cos(angle) - tmpV.Y*Sin(angle)
	//b.y = a.x * sin(angle) + a.y * cos(angle);
	v.Y = tmpV.X*Sin(angle) + tmpV.Y*Cos(angle)
}
