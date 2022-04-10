package main

import (
	"3d/internal/objects"
	"3d/internal/vector"
)

var (
	pixelAspect      = 7.0 / 16.0
	gradientSymbols1 = " .:!/|(41lZH9W8$@"
	gradientSymbols2 = Reverse("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ")
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	light := vector.NewVec3(-0.5, 0.5, -1.0)
	sphere := objects.NewSphere(vector.NewVec3(0, 3, 0), 1)
	box := objects.NewBox(vector.NewVec3(0, 0, -0.1), vector.NewVec3(1))
	plane := objects.NewPlane(vector.NewVec3(0, 0, -1))
	screen, err := objects.NewScreen(pixelAspect, gradientSymbols1)
	if err != nil {
		panic(err)
	}

	for t := 0.0; t < 12600.0; t++ {
		for i := uint64(0); i < screen.Width; i++ {
			for j := uint64(0); j < screen.Height; j++ {
				uv := vector.NewVec2(float64(i), float64(j)).Div(vector.NewVec2(float64(screen.Width), float64(screen.Height))).Mult(2.0).Diff(1.0)
				uv.X *= screen.Aspect
				camera := objects.NewCamera(vector.NewVec3(-6, 0, 0), vector.NewVec3(2, uv.X, uv.Y))
				camera.Position.RotateY(0.25)
				camera.Direction.RotateY(0.25)
				camera.Position.RotateZ(t * 0.01)
				camera.Direction.RotateZ(t * 0.01)
				diff := 1.0
				for k := 0; k < 5; k++ {
					minIt := 99999.0
					albedo := 1.0
					intersection := sphere.Intersection(camera)
					n := vector.NewVec3(0)
					if intersection.X > 0 {
						itPoint := camera.Direction.Mult(intersection.X).Sum(camera.Position.Diff(sphere.Position))
						minIt = intersection.X
						n = itPoint.Norm()
					}
					intersection, boxNorm := box.Intersection(camera)
					if intersection.X > 0 && intersection.X < minIt {
						minIt = intersection.X
						n = boxNorm
					}
					intersection = plane.Intersection(camera)
					if intersection.X > 0 && intersection.X < minIt {
						minIt = intersection.X
						n = vector.NewVec3(0, 0, -1)
						albedo = 0.5
					}
					if minIt < 99999 {
						diff *= (n.Dot(light)*0.5 + 0.5) * albedo
						camera.Position = camera.Position.Sum(camera.Direction.Mult(minIt - 0.01))
						camera.Direction = camera.Direction.Reflect(n)
					} else {
						break
					}
				}
				color := int(diff * float64(len(screen.Gradient)))
				color = vector.Clamp(color, 0, len(screen.Gradient)-1)
				pixel := screen.Gradient[color]
				screen.Screen[i+j*screen.Width] = pixel
			}
		}
		screen.Render()
	}
}
