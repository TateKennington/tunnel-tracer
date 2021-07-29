package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/TateKennington/tunnel-tracer/geometry"
)

func main() {
	const aspect_ratio = 16.0 / 9.0
	const image_width = 400.0
	const image_height = image_width / aspect_ratio

	image_rect := image.Rect(0, 0, image_width, image_height)
	render := image.NewRGBA(image_rect)

	const focal_length = 1.0
	const viewport_height = 2.0
	const viewport_width = viewport_height * aspect_ratio

	origin := geometry.Vec3{0, 0, 0}
	horizontal := geometry.Vec3{viewport_width, 0, 0}
	vertical := geometry.Vec3{0, -viewport_height, 0}
	viewport_origin := geometry.Vec3{-viewport_width / 2, viewport_height / 2, -focal_length}

	for y := 0; y < image_height; y++ {
		for x := 0; x < image_width; x++ {
			fmt.Printf("\rProgress %d/%.0f", y*image_width+x+1, image_width*image_height)
			u := float64(x) / (image_width - 1)
			v := float64(y) / (image_height - 1)
			direction := origin.LineTo(
				viewport_origin.Translate(
					horizontal.Scale(u),
				).Translate(
					vertical.Scale(v),
				),
			)
			r := geometry.Ray{
				Origin:    origin,
				Direction: direction,
			}
			render_color := ray_color(r)
			render.Set(x, y, render_color)
		}
	}

	output, err := os.Create("dist/output.png")
	if err != nil {
		fmt.Printf("Error opening output file: %s", err.Error())
	}

	png.Encode(output, render)
}

func ray_color(r geometry.Ray) color.Color {
	unit_direction := r.Direction.Unit()
	t := 0.5 * (unit_direction.Y + 1.0)
	return geometry.Vec3{1.0, 1.0, 1.0}.Lerp(t, geometry.Vec3{0.5, 0.7, 1.0})
}
