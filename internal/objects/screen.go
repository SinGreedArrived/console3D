package objects

import (
	"fmt"

	"3d/internal/terminal"
)

type Screen struct {
	Width, Height uint64
	Aspect        float64
	Gradient      []rune
	Screen        []rune
}

func NewScreen(pixelAspect float64, gradient string) (*Screen, error) {
	width, height, err := terminal.GetResolution()
	if err != nil {
		fmt.Println(fmt.Errorf("GetResolution: %w", err))
		return nil, fmt.Errorf("terminal.GetResolution: %w", err)
	}
	return &Screen{
		Width:    width,
		Height:   height,
		Aspect:   float64(width) / float64(height) * pixelAspect,
		Gradient: []rune(gradient),
		Screen:   make([]rune, width*height),
	}, nil
}

func (s *Screen) Render() {
	fmt.Printf("\x1B[H%s", string(s.Screen))
	// fmt.Printf("%s", string(s.Screen))
}

func (s *Screen) SetGradient(gradient string) {
	s.Gradient = []rune(gradient)
}
