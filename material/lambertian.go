package material

import "github.com/TateKennington/tunnel-tracer/geometry"

type Lambertian struct {
	Albedo geometry.Vec3
}

//Implement Material for Lambertian
func (lambertian Lambertian) Scatter(ray geometry.Ray, hit geometry.HitResult) (bool, geometry.Vec3, geometry.Ray) {
	rayDirection := geometry.RandomVecSphere().Translate(hit.Normal)
	if rayDirection.NearZero() {
		rayDirection = hit.Normal
	}
	diffuseRay := geometry.Ray{
		Origin:    hit.Point,
		Direction: rayDirection,
	}
	return true, lambertian.Albedo, diffuseRay
}
