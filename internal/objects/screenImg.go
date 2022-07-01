package objects

import (
	"fmt"

	"3d/internal/terminal"
	"3d/internal/vector"
)

type PicScreen struct {
	Width, Height uint64
	Aspect        float64
	Gradient      []rune
	Screen        []rune
	Coord         *vector.Vec2
}

func NewScreenImg(pixelAspect float64, gradient string) (Screen, error) {
	width, height, err := terminal.GetResolution()
	if err != nil {
		fmt.Println(fmt.Errorf("GetResolution: %w", err))
		return nil, fmt.Errorf("terminal.GetResolution: %w", err)
	}
	return &PicScreen{
		Width:    width,
		Height:   height,
		Aspect:   float64(width) / float64(height) * pixelAspect,
		Gradient: []rune(gradient),
		Screen:   make([]rune, width*height),
		Coord:    vector.NewVec2(0., -1.),
	}, nil
}

func (s *PicScreen) GetSize() (uint64, uint64) {
	return 0, 0
}

func (s *PicScreen) Render() {
	fmt.Printf("\x1B[H%s", string(s.Screen))
}

func (s *PicScreen) SetGradient(gradient string) {
	s.Gradient = []rune(gradient)
}

func (s *PicScreen) NextCoord() bool {
	s.Coord.Y += 1.
	if s.Coord.X == float64(s.Height)-1. && s.Coord.Y == float64(s.Width)-1. {
		s.Coord.X = 0
		s.Coord.Y = 0
		return false
	}
	if s.Coord.Y == float64(s.Width) {
		s.Coord.Y = 0
		s.Coord.X += 1
	}

	return true
}

func (s *PicScreen) UV(coord *vector.Vec2) *vector.Vec2 {
	uv := vector.NewVec2(s.Coord.Y, s.Coord.X).Div(vector.NewVec2(float64(s.Width), float64(s.Height))).Mult(2.0).Diff(1.0)
	uv.X *= s.Aspect
	return uv
}

func (s *PicScreen) SetPixel(diff float64, coord *vector.Vec2) {
	panic("not impl")
}

func (s *PicScreen) GetCoord() *vector.Vec2 {
	return s.Coord
}
