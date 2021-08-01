package scene

import "github.com/TateKennington/tunnel-tracer/geometry"

type Camera struct {
	origin          geometry.Vec3
	horizontal      geometry.Vec3
	vertical        geometry.Vec3
	viewport_origin geometry.Vec3
}

func NewCamera(aspect_ratio float64) Camera {
	const focal_length = 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * aspect_ratio

	origin := geometry.Vec3{0, 0, 0}
	horizontal := geometry.Vec3{viewport_width, 0, 0}
	vertical := geometry.Vec3{0, -viewport_height, 0}
	viewport_origin := geometry.Vec3{-viewport_width / 2, viewport_height / 2, -focal_length}

	return Camera{
		origin,
		horizontal,
		vertical,
		viewport_origin,
	}
}

func (camera *Camera) GetRay(u, v float64) geometry.Ray {
	direction := camera.origin.LineTo(
		camera.viewport_origin.Translate(
			camera.horizontal.Scale(u),
		).Translate(
			camera.vertical.Scale(v),
		),
	)
	return geometry.Ray{
		Origin:    camera.origin,
		Direction: direction,
	}
}
