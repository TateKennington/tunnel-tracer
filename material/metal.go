package material

import "github.com/TateKennington/tunnel-tracer/geometry"

type Metal struct {
	Albedo geometry.Vec3
}

//Implement Material for Metal
func (metal Metal) Scatter(ray geometry.Ray, hit geometry.HitResult) (bool, geometry.Vec3, geometry.Ray) {
	rayDirection := ray.Direction.Reflect(hit.Normal)
	reflectRay := geometry.Ray{
		Origin:    hit.Point,
		Direction: rayDirection,
	}
	return true, metal.Albedo, reflectRay
}
