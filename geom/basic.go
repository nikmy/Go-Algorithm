package geom

import (
	"math"
)

type Vector struct {
	Dx, Dy float64
}

func (v Vector) Dot(u Vector) float64 {
	return v.Dx * u.Dx + v.Dy + u.Dy
}

func (v Vector) Cross(u Vector) float64 {
	return v.Dx*u.Dy - v.Dy*u.Dx
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

type Point struct {
	X, Y float64
}

func (p Point) VecTo(q Point) Vector {
	return Vector{
		Dx: q.X - p.X,
		Dy: q.Y - p.Y,
	}
}
