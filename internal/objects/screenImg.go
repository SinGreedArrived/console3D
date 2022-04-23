package objects

import (
	"fmt"

	"3d/internal/terminal"
	"3d/internal/vector"
)

type ScreenImg struct {
	Width, Height uint64
	Aspect        float64
	Gradient      []rune
	Screen        []rune
	Coord         *vector.Vec2
}

func NewScreenImg(pixelAspect float64, gradient string) (IScreen, error) {
	width, height, err := terminal.GetResolution()
	if err != nil {
		fmt.Println(fmt.Errorf("GetResolution: %w", err))
		return nil, fmt.Errorf("terminal.GetResolution: %w", err)
	}
	return &ScreenImg{
		Width:    width,
		Height:   height,
		Aspect:   float64(width) / float64(height) * pixelAspect,
		Gradient: []rune(gradient),
		Screen:   make([]rune, width*height),
		Coord:    vector.NewVec2(0., -1.),
	}, nil
}

func (s *ScreenImg) Render() {
	fmt.Printf("\x1B[H%s", string(s.Screen))
}

func (s *ScreenImg) SetGradient(gradient string) {
	s.Gradient = []rune(gradient)
}

func (s *ScreenImg) NextCoord() bool {
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

func (s *ScreenImg) UV() *vector.Vec2 {
	uv := vector.NewVec2(s.Coord.Y, s.Coord.X).Div(vector.NewVec2(float64(s.Width), float64(s.Height))).Mult(2.0).Diff(1.0)
	uv.X *= s.Aspect
	return uv
}

func (s *ScreenImg) SetPixel(diff float64) {
	color := int(diff * float64(len(s.Gradient)))
	color = vector.Clamp(color, 0, len(s.Gradient)-1)
	pixel := s.Gradient[color]
	s.Screen[uint64(s.Coord.Y)+uint64(s.Coord.X)*s.Width] = pixel
}
