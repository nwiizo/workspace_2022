package main

import (
	"image/png"
	"log"
	"os"

	"github.com/esimov/triangle"
)

func main() {

	proc := &triangle.Processor{
		MaxPoints:  500,
		BlurRadius: 2,
		PointRate:  1,
		BlurFactor: 1,
		EdgeFactor: 6,
	}

	img := &triangle.Image{
		Processor: *proc,
	}

	input, err := os.Open("sample/input.jpeg")
	if err != nil {
		log.Fatalf("error opening the source file: %v", err)
	}

	// decode image
	src, err := img.DecodeImage(input)
	if err != nil {
		log.Fatalf("error decoding the image: %v", err)
	}
	res, _, _, err := img.Draw(src, *proc, func() {})
	if err != nil {
		log.Fatalf("error generating the triangles: %v", err)
	}

	output, err := os.Create("sample/output.jpeg")
	if err != nil {
		log.Fatalf("error opening the destination file: %v", err)
	}

	// encode image
	png.Encode(output, res)
}
