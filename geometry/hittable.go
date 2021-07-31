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

type Scene struct {
	objects []Hittable
}

func (scene *Scene) Add(object Hittable) {
	scene.objects = append(scene.objects, object)
}

//Implement Hittable for scene
func (scene Scene) Hit(ray Ray, minT, maxT float64) (bool, HitResult) {
	hit := false
	result := HitResult{}
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
