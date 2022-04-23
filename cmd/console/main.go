package main

import (
	"3d/internal/helper"
	"3d/internal/objects"
	"3d/internal/vector"
)

var (
	pixelAspect      = 7.0 / 16.0
	gradientSymbols1 = " .:!/|(41lZH9W8$@"
	gradientSymbols2 = helper.Reverse("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ")
)

func main() {
	light := vector.NewVec3(
		-0.5,
		0.5,
		-1.0,
	)
	sphere := objects.NewSphere(
		vector.NewVec3(0, 3, 0),
		1,
	)
	box := objects.NewBox(
		vector.NewVec3(0, 0, -0.1),
		vector.NewVec3(1),
	)
	plane := objects.NewPlane(
		vector.NewVec3(0, 0, -1),
	)
	screen, err := objects.NewScreen(
		pixelAspect,
		gradientSymbols1,
	)
	if err != nil {
		panic(err)
	}

	for t := 0.0; t < 12600.0; t++ {
		for screen.NextCoord() {
			camera := objects.NewCamera(
				vector.NewVec3(-6, 0, 0),
				vector.NewVec3(2, screen.UV()),
			)
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
			screen.SetPixel(diff)
		}
		screen.Render()
	}
}
