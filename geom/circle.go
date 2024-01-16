package geom

type Circle struct {
	O Point
	R float64
}

func (c Circle) Intersect(d Circle) bool {
	return c.O.VecTo(d.O).Length() < c.R+d.R
}

func (c Circle) Encloses(p Point) bool {
	return c.O.VecTo(p).Length() < c.R
}
