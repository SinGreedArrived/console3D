package main

import (
	"3d/internal/terminal"
	"3d/internal/vector"
	"fmt"
	"os"
)

func reverseRune(input []rune) []rune {
	if len(input) == 0 {
		return input
	}
	return append(reverseRune(input[1:]), input[0])
}

func main() {
	width, height, err := terminal.GetResolution()
	if err != nil {
		fmt.Println(fmt.Errorf("GetResolution: %w", err))
		os.Exit(1)
	}
	aspect := float64(width) / float64(height)
	pixelAspect := 7.0 / 16.0
	gradient := []rune(" .:!/|(41lZH9W8$@")
	screen := make([]rune, width*height)
	for t := 0.0; t < 12600.0; t++ {
		light := vector.NewVec3(-0.5, 0.5, -1.0)
		spherePos := vector.NewVec3(0, 3, 0)
		boxPos := vector.NewVec3(0, 0, -0.1)
		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				uv := vector.NewVec2(float64(i), float64(j)).Div(vector.NewVec2(float64(width), float64(height))).Mult(2.0).Diff(1.0)
				uv.X *= aspect * pixelAspect
				ro := vector.NewVec3(-6, 0, 0)
				rd := vector.NewVec3(2, uv).Norm()
				ro = vector.RotateY(ro, 0.25)
				rd = vector.RotateY(rd, 0.25)
				ro = vector.RotateZ(ro, t*0.01)
				rd = vector.RotateZ(rd, t*0.01)
				diff := 1.0
				for k := 0; k < 5; k++ {
					minIt := 99999.0
					albedo := 1.0
					intersection := vector.Sphere(ro.Diff(spherePos), rd, 1)
					n := vector.NewVec3(0)
					if intersection.X > 0 {
						itPoint := rd.Mult(intersection.X).Sum(ro.Diff(spherePos))
						minIt = intersection.X
						n = itPoint.Norm()
					}
					boxN := vector.NewVec3(0)
					intersection = vector.Box(ro.Diff(boxPos), rd, vector.NewVec3(1.0), &boxN)
					if intersection.X > 0 && intersection.X < minIt {
						minIt = intersection.X
						n = boxN
					}
					intersection = vector.NewVec2(vector.Plane(ro, rd, vector.NewVec3(0, 0, -1), 1))
					if intersection.X > 0 && intersection.X < minIt {
						minIt = intersection.X
						n = vector.NewVec3(0, 0, -1)
						albedo = 0.5
					}
					if minIt < 99999 {
						diff *= (vector.Dot3(n, light)*0.5 + 0.5) * albedo
						ro = ro.Sum(rd.Mult(minIt - 0.01))
						rd = vector.Reflect3(rd, n)
					} else {
						break
					}
				}
				color := int(diff * float64(len(gradient)))
				color = vector.Clamp(color, 0, len(gradient)-1)
				pixel := gradient[color]
				screen[i+j*width] = pixel
			}
		}
		fmt.Print(string(screen))
	}
}
