package scene

import "github.com/TateKennington/tunnel-tracer/geometry"

type Scene struct {
	objects []Object
}

func (scene *Scene) Add(object Object) {
	scene.objects = append(scene.objects, object)
}

func (scene Scene) Hit(ray geometry.Ray, minT, maxT float64) (bool, ExtendedHitResult) {
	hit := false
	result := ExtendedHitResult{}
	bestT := maxT
	for _, object := range scene.objects {
		if newHit, newResult := object.Hit(ray, minT, bestT); newHit {
			result = newResult
			hit = true
			bestT = newResult.T
		}
	}
	return hit, result
}
