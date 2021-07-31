package geometry

import (
	"math"
)

type Sphere struct {
	Origin Vec3
	Radius float64
}

//Implement Hittable for Sphere
func (sphere Sphere) Hit(ray Ray, minT, maxT float64) (bool, HitResult) {
	sphereToRay := sphere.Origin.LineTo(ray.Origin)
	halfB := ray.Direction.Dot(sphereToRay)
	directionSq := ray.Direction.LenSq()
	determinant := halfB*halfB - directionSq*(sphereToRay.LenSq()-sphere.Radius*sphere.Radius)

	if determinant < 0 {
		return false, HitResult{}
	}

	root := (-halfB - math.Sqrt(determinant)) / directionSq
	if root < minT || root > maxT {
		root = (-halfB + math.Sqrt(determinant)) / directionSq
		if root < minT || root > maxT {
			return false, HitResult{}
		}
	}

	point := ray.At(root)
	normal := sphere.Origin.LineTo(point).Unit()
	frontFace := ray.Direction.Dot(normal) <= 0.0

	if !frontFace {
		normal = normal.Scale(-1)
	}

	return true, HitResult{
		Point:     point,
		T:         root,
		Normal:    normal,
		FrontFace: frontFace,
	}
}
