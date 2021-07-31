package geometry

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (ray *Ray) At(t float64) Vec3 {
	return ray.Origin.Translate(ray.Direction.Scale(t))
}
