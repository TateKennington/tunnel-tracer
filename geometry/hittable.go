package geometry

type HitResult struct {
	Point     Vec3
	T         float64
	Normal    Vec3
	FrontFace bool
}

type Hittable interface {
	Hit(ray Ray, minT, maxT float64) (bool, HitResult)
}
