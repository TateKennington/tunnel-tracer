package geometry

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (self Vec3) Translate(other Vec3) Vec3 {
	self.X += other.X
	self.Y += other.Y
	self.Z += other.Z
	return self
}

func (self Vec3) Scale(factor float64) Vec3 {
	self.X *= factor
	self.Y *= factor
	self.Z *= factor
	return self
}

func (self Vec3) LineTo(other Vec3) Vec3 {
	return other.Translate(self.Scale(-1))
}

func (self Vec3) LenSq() float64 {
	return self.X*self.X + self.Y*self.Y + self.Z*self.Z
}

func (self Vec3) Len() float64 {
	return math.Sqrt(self.LenSq())
}

func (self Vec3) Dot(other Vec3) float64 {
	return self.X*other.X + self.Y*other.Y + self.Z*other.Z
}

func (self Vec3) Cross(other Vec3) Vec3 {
	X := self.Y*other.Z - self.Z*other.Y
	Y := self.Z*other.X - self.X*other.Z
	Z := self.X*other.Y - self.Y*other.X
	return Vec3{X, Y, Z}
}

func (self Vec3) Unit() Vec3 {
	return self.Scale(1 / self.Len())
}

func (self Vec3) Lerp(t float64, other Vec3) Vec3 {
	return self.Scale(1.0 - t).Translate(other.Scale(t))
}

//Implement Stringer
func (self Vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", self.X, self.Y, self.Z)
}

//Implement Color
func (self Vec3) RGBA() (r, g, b, a uint32) {
	var max uint32 = 0xFFFF
	self = self.Scale(float64(max))
	return uint32(self.X), uint32(self.Y), uint32(self.Z), max
}
