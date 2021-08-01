package scene

import (
	"github.com/TateKennington/tunnel-tracer/geometry"
	"github.com/TateKennington/tunnel-tracer/material"
)

type Object struct {
	Object   geometry.Hittable
	Material material.Material
}

type ExtendedHitResult struct {
	material.Material
	geometry.HitResult
}

func (object *Object) Hit(ray geometry.Ray, minT, maxT float64) (bool, ExtendedHitResult) {
	hit, result := object.Object.Hit(ray, minT, maxT)
	return hit, ExtendedHitResult{
		HitResult: result,
		Material:  object.Material,
	}
}
