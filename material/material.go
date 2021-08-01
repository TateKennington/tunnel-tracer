package material

import "github.com/TateKennington/tunnel-tracer/geometry"

type Material interface {
	Scatter(ray geometry.Ray, hit geometry.HitResult) (bool, geometry.Vec3, geometry.Ray)
}
