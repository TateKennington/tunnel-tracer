package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"

	"github.com/TateKennington/tunnel-tracer/geometry"
	"github.com/TateKennington/tunnel-tracer/material"
	"github.com/TateKennington/tunnel-tracer/scene"
)

func main() {
	const aspect_ratio = 16.0 / 9.0
	const image_width = 400.0
	const image_height = image_width / aspect_ratio
	const samples_per_pixel = 100
	const max_depth = 50

	image_rect := image.Rect(0, 0, image_width, image_height)
	render := image.NewRGBA(image_rect)

	camera := scene.NewCamera(aspect_ratio)

	material_ground := material.Lambertian{geometry.Vec3{0.8, 0.8, 0.0}}
	material_center := material.Lambertian{geometry.Vec3{0.7, 0.3, 0.3}}
	material_left := material.Metal{geometry.Vec3{0.8, 0.8, 0.8}}
	material_right := material.Metal{geometry.Vec3{0.8, 0.6, 0.2}}

	world := scene.Scene{}
	world.Add(scene.Object{geometry.Sphere{geometry.Vec3{0.0, -100.5, -1.0}, 100.0}, material_ground})
	world.Add(scene.Object{geometry.Sphere{geometry.Vec3{0.0, 0.0, -1.0}, 0.5}, material_center})
	world.Add(scene.Object{geometry.Sphere{geometry.Vec3{-1.0, 0.0, -1.0}, 0.5}, material_left})
	world.Add(scene.Object{geometry.Sphere{geometry.Vec3{1.0, 0.0, -1.0}, 0.5}, material_right})

	for y := 0; y < image_height; y++ {
		fmt.Printf("\rProgress %d/%.0f", y+1, image_height)
		for x := 0; x < image_width; x++ {
			pixelColor := geometry.Vec3{0, 0, 0}
			for s := 0; s < samples_per_pixel; s++ {
				u := (float64(x) + rand.Float64()) / (image_width - 1)
				v := (float64(y) + rand.Float64()) / (image_height - 1)
				r := camera.GetRay(u, v)
				pixelColor.Add(ray_color(r, world, max_depth))
			}
			pixelColor.Mult(1.0 / float64(samples_per_pixel))
			pixelColor.Sqrt()
			render.Set(x, y, pixelColor)
		}
	}

	output, err := os.Create("dist/output.png")
	if err != nil {
		fmt.Printf("Error opening output file: %s", err.Error())
	}

	png.Encode(output, render)
}

func ray_color(r geometry.Ray, scene scene.Scene, depth int) geometry.Vec3 {
	if depth <= 0 {
		return geometry.Vec3{0, 0, 0}
	}
	if hit, result := scene.Hit(r, 0.001, math.MaxFloat64); hit {
		scatter, attenuation, diffuseRay := result.Material.Scatter(r, result.HitResult)
		if scatter {
			color := ray_color(diffuseRay, scene, depth-1)
			color.VecMult(attenuation)
			return color
		}
		return geometry.Vec3{0, 0, 0}
	}
	unit_direction := r.Direction.Unit()
	t := 0.5 * (unit_direction.Y + 1.0)
	return geometry.Vec3{1.0, 1.0, 1.0}.Lerp(t, geometry.Vec3{0.5, 0.7, 1.0})
}
