package vector

import . "math"

type Vec2 struct {
	X float64
	Y float64
}

func NewVec2(i ...interface{}) *Vec2 {
	return newVec2(i)
}

func newVec2(i []interface{}) *Vec2 {
	v := &Vec2{}
	if len(i) == 1 {
		switch i[0].(type) {
		case int:
			return &Vec2{
				X: float64(i[0].(int)),
				Y: float64(i[0].(int)),
			}
		case float64:
			return &Vec2{
				X: float64(i[0].(float64)),
				Y: float64(i[0].(float64)),
			}
		case *Vec3:
			return &Vec2{
				X: i[0].(*Vec3).X,
				Y: i[0].(*Vec3).Y,
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
			}
		case uint64:
			switch n {
			case 0:
				v.X = float64(i[n].(uint64))
			case 1:
				v.Y = float64(i[n].(uint64))
			}
		case float64:
			switch n {
			case 0:
				v.X = i[n].(float64)
			case 1:
				v.Y = i[n].(float64)
			}
		case *Vec2:
			switch n {
			case 0:
				v.X = i[n].(*Vec2).X
				v.Y = i[n].(*Vec2).Y
			}
		}
	}
	return v
}

func (v *Vec2) Sum(i interface{}) *Vec2 {
	switch i.(type) {
	case *Vec2:
		vector2 := i.(*Vec2)
		return &Vec2{
			X: v.X + vector2.X,
			Y: v.Y + vector2.Y,
		}
	case float64:
		f := i.(float64)
		return &Vec2{
			X: v.X + f,
			Y: v.Y + f,
		}
	}
	return nil
}

func (v *Vec2) Diff(i interface{}) *Vec2 {
	switch i.(type) {
	case *Vec2:
		vector2 := i.(*Vec2)
		return &Vec2{
			X: v.X - vector2.X,
			Y: v.Y - vector2.Y,
		}
	case float64:
		f := i.(float64)
		return &Vec2{
			X: v.X - f,
			Y: v.Y - f,
		}
	}
	return nil
}

func (v *Vec2) Mult(i interface{}) *Vec2 {
	switch i.(type) {
	case *Vec2:
		vector2 := i.(*Vec2)
		return &Vec2{
			X: v.X * vector2.X,
			Y: v.Y * vector2.Y,
		}
	case float64:
		f := i.(float64)
		return &Vec2{
			X: v.X * f,
			Y: v.Y * f,
		}
	}
	return nil
}

func (v *Vec2) Div(i interface{}) *Vec2 {
	switch i.(type) {
	case *Vec2:
		vector2 := i.(*Vec2)
		return &Vec2{
			X: v.X / vector2.X,
			Y: v.Y / vector2.Y,
		}
	case float64:
		f := i.(float64)
		return &Vec2{
			X: v.X / f,
			Y: v.Y / f,
		}
	}
	return nil
}

func (v *Vec2) Len() float64 {
	return Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) Norm() *Vec2 {
	return v.Div(v.Len())
}
