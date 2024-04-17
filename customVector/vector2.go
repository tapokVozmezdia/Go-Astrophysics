package customVector

import (
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

func VectorAdd(v1 Vector2, v2 Vector2) Vector2 {
	return Vector2{v1.X + v2.X, v1.Y + v2.Y}
}

// format: (SOURCE, TARGET)
func VectorDiff(v1 Vector2, v2 Vector2) Vector2 {
	return Vector2{v2.X - v1.X, v2.Y - v1.Y}
}

func (vec Vector2) GetLen() float64 {
	len_squared := vec.X*vec.X + vec.Y*vec.Y
	return math.Sqrt(len_squared)
}

// returns vector copy with multiplied length
func (vec Vector2) GetMultiplied(coef float64) Vector2 {
	v := Vector2{vec.X * coef, vec.Y * coef}
	return v
}

// reverses vector
func (vec *Vector2) Reverse() {
	vec.X *= -1
	vec.Y *= -1
}

// returns reversed vector copy
func (vec *Vector2) GetReversed() Vector2 {
	v := Vector2{vec.X * -1, vec.Y * -1}
	return v
}

// Multiplies vector's length by an argument
func (vec *Vector2) MultLen(coef float64) {
	vec.X *= coef
	vec.Y *= coef
}

func (vec *Vector2) Normalize() {
	len := vec.GetLen()
	vec.X /= len
	vec.Y /= len
}

func GetLenBetweenVectors(v1 *Vector2, v2 *Vector2) float64 {
	v3 := Vector2{v1.X - v2.X, v1.Y - v2.Y}
	return v3.GetLen()
}
