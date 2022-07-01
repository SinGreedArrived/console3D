package objects

import (
	"fmt"

	"3d/internal/terminal"
	"3d/internal/vector"
)

type Frame struct {
	Width, Height uint64
	Aspect        float64
	Gradient      []rune
	Screen        []rune
	Coord         *vector.Vec2
}

func NewFrame(pixelAspect float64, gradient string) (*Frame, error) {
	width, height, err := terminal.GetResolution()
	if err != nil {
		fmt.Println(fmt.Errorf("GetResolution: %w", err))
		return nil, fmt.Errorf("terminal.GetResolution: %w", err)
	}
	return &Frame{
		Width:    width,
		Height:   height,
		Aspect:   float64(width) / float64(height) * pixelAspect,
		Gradient: []rune(gradient),
		Screen:   make([]rune, width*height),
		Coord:    vector.NewVec2(0., -1.),
	}, nil
}

func (s *Frame) GetSize() (uint64, uint64) {
	return s.Width, s.Height
}

func (s *Frame) Render() {
	fmt.Printf("\x1B[H%s", string(s.Screen))
}

func (s *Frame) SetGradient(gradient string) {
	s.Gradient = []rune(gradient)
}

func (s *Frame) NextCoord() bool {
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

func (s *Frame) UV(coord *vector.Vec2) *vector.Vec2 {
	if coord == nil {
		uv := vector.NewVec2(s.Coord.Y, s.Coord.X).Div(vector.NewVec2(float64(s.Width), float64(s.Height))).Mult(2.0).Diff(1.0)
		uv.X *= s.Aspect
		return uv
	}
	uv := vector.NewVec2(coord.Y, coord.X).Div(vector.NewVec2(float64(s.Width), float64(s.Height))).Mult(2.0).Diff(1.0)
	uv.X *= s.Aspect
	return uv
}

func (s *Frame) GetCoord() *vector.Vec2 {
	return s.Coord
}

func (s *Frame) SetPixel(diff float64, coord *vector.Vec2) {
	color := int(diff * float64(len(s.Gradient)))
	color = vector.Clamp(color, 0, len(s.Gradient)-1)
	pixel := s.Gradient[color]
	if coord == nil {
		s.Screen[uint64(s.Coord.Y)+uint64(s.Coord.X)*s.Width] = pixel
		return
	}
	s.Screen[uint64(coord.Y)+uint64(coord.X)*s.Width] = pixel
}
