package geometry

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (self *Ray) At(t float64) Vec3 {
	return self.Origin.Translate(self.Direction.Scale(t))
}
