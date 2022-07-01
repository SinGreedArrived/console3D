package main

import (
	"fmt"
	"sync"
	"time"

	"3d/internal/objects"
	"3d/internal/vector"
)

var (
	pixelAspect      = 7.0 / 16.0
	gradientSymbols1 = " .:!/|(41lZH9W8$@"
)

func createFrame(
	number float64,
	light vector.Vec3,
	sphere objects.Sphere,
	box objects.Box,
	plane objects.Plane,
) (*objects.Frame, error) {
	frame, err := objects.NewFrame(
		pixelAspect,
		gradientSymbols1,
	)
	if err != nil {
		return nil, fmt.Errorf("objects.NewFrame: %w", err)
	}

	w, h := frame.GetSize()

	for i := uint64(0); i < h; i++ {
		for j := uint64(0); j < w; j++ {
			xy := vector.NewVec2(i, j)
			camera := objects.NewCamera(
				vector.NewVec3(-6, 0, 0),
				vector.NewVec3(2, frame.UV(xy)),
			)
			camera.Position.RotateY(0.25)
			camera.Direction.RotateY(0.25)
			camera.Position.RotateZ(number * 0.01)
			camera.Direction.RotateZ(number * 0.01)
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
			frame.SetPixel(diff, xy)
		}
	}
	return frame, nil
}

type task struct {
	number uint64
	light  vector.Vec3
	sphere objects.Sphere
	box    objects.Box
	plane  objects.Plane
}

const (
	delayFrame   = time.Duration(20)
	requestFrame = time.Duration(10)
)

func main() {
	var wg sync.WaitGroup
	workerPoolSize := 8

	dataCh := make(chan task, workerPoolSize)
	synFrame := make(chan uint64)
	render := make(chan *objects.Frame)

	for i := 0; i < workerPoolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for data := range dataCh {
				frame, _ := createFrame(float64(data.number), data.light, data.sphere, data.box, data.plane)
				for {
					need := <-synFrame
					if need == data.number {
						render <- frame
						break
					} else {
						synFrame <- need
						time.Sleep(time.Microsecond * requestFrame)
					}
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		number := uint64(0)
		defer wg.Done()

		synFrame <- number
		for frame := range render {
			frame.Render()
			number = number + 1
			synFrame <- number
			time.Sleep(time.Millisecond * delayFrame)
		}
	}()

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

	for t := uint64(0); t < 12600; t++ {
		dataCh <- task{
			number: t,
			light:  *light,
			sphere: *sphere,
			box:    *box,
			plane:  *plane,
		}
	}
}
