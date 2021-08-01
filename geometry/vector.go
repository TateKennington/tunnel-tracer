package geometry

import (
	"fmt"
	"math"
	"math/rand"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (vec Vec3) Translate(other Vec3) Vec3 {
	vec.X += other.X
	vec.Y += other.Y
	vec.Z += other.Z
	return vec
}

func (vec *Vec3) Add(other Vec3) {
	vec.X += other.X
	vec.Y += other.Y
	vec.Z += other.Z
}

func (vec *Vec3) VecMult(other Vec3) {
	vec.X *= other.X
	vec.Y *= other.Y
	vec.Z *= other.Z
}

func (vec *Vec3) Mult(factor float64) {
	vec.X *= factor
	vec.Y *= factor
	vec.Z *= factor
}

func (vec *Vec3) Sqrt() {
	vec.X = math.Sqrt(vec.X)
	vec.Y = math.Sqrt(vec.Y)
	vec.Z = math.Sqrt(vec.Z)
}

func (vec Vec3) Scale(factor float64) Vec3 {
	vec.X *= factor
	vec.Y *= factor
	vec.Z *= factor
	return vec
}

func (vec Vec3) Reflect(other Vec3) Vec3 {
	return vec.Translate(other.Scale(-2 * vec.Dot(other)))
}

func (vec Vec3) LineTo(other Vec3) Vec3 {
	return other.Translate(vec.Scale(-1))
}

func (vec Vec3) LenSq() float64 {
	return vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z
}

func (vec Vec3) Len() float64 {
	return math.Sqrt(vec.LenSq())
}

func (vec Vec3) Dot(other Vec3) float64 {
	return vec.X*other.X + vec.Y*other.Y + vec.Z*other.Z
}

func (vec Vec3) Cross(other Vec3) Vec3 {
	X := vec.Y*other.Z - vec.Z*other.Y
	Y := vec.Z*other.X - vec.X*other.Z
	Z := vec.X*other.Y - vec.Y*other.X
	return Vec3{X, Y, Z}
}

func (vec Vec3) Unit() Vec3 {
	return vec.Scale(1 / vec.Len())
}

func (vec Vec3) Lerp(t float64, other Vec3) Vec3 {
	return vec.Scale(1.0 - t).Translate(other.Scale(t))
}

func (vec *Vec3) NearZero() bool {
	const e = 1e-8
	return math.Abs(vec.X) < e && math.Abs(vec.Y) < e && math.Abs(vec.Z) < e
}

func RandomVec(min, max float64) Vec3 {
	return Vec3{
		X: min + rand.Float64()*(max-min),
		Y: min + rand.Float64()*(max-min),
		Z: min + rand.Float64()*(max-min),
	}
}

func RandomVecSphere() Vec3 {
	for vec := RandomVec(-1, 1); ; vec = RandomVec(-1, 1) {
		if vec.LenSq() <= 1 {
			return vec.Unit()
		}
	}
}

//Implement Stringer
func (vec Vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", vec.X, vec.Y, vec.Z)
}

//Implement Color
func (vec Vec3) RGBA() (r, g, b, a uint32) {
	var max uint32 = 0xFFFF
	vec = vec.Scale(float64(max))
	return uint32(vec.X), uint32(vec.Y), uint32(vec.Z), max
}
