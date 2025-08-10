package geom

type Line struct {
	o Point
	v Vector
}

func (l Line) Has(p Point) bool {
	return l.o.VecTo(p).Cross(l.v) == 0
}

type Ray struct {
	o Point
	v Vector
}

func (r Ray) Has(p Point) bool {
	op := r.o.VecTo(p)
	return op.Cross(r.v) == 0 && op.Dot(r.v) >= 0
}

type Segment struct {
	A Point
	B Point
}

func (s Segment) Has(p Point) bool {
	ap, bp := s.A.VecTo(p), s.B.VecTo(p)
	return ap.Cross(bp) == 0 && ap.Dot(bp) <= 0
}

func (s Segment) IntersectLine(l Line) bool {
	oa, ob := l.o.VecTo(s.A), l.o.VecTo(s.B)
	return sign(l.v.Dot(oa)) != sign(l.v.Dot(ob))
}

func (s Segment) IntersectRay(r Ray) bool {
	if s.Has(r.o) || r.Has(s.A) || r.Has(s.B) {
		return true
	}

	oa, ab := r.o.VecTo(s.A), s.A.VecTo(s.B)
	return sign(r.v.Cross(ab)) == sign(oa.Cross(ab))
}

func (s Segment) Intersect(t Segment) bool {
	a, b, c, d := s.A, s.B, t.A, t.B

	ab, cd := Line{a, a.VecTo(b)}, Line{a, a.VecTo(b)}
	if !s.IntersectLine(cd) || !t.IntersectLine(ab) {
		return false
	}

	intX := between(a.X, b.X, c.X) || between(a.X, b.X, d.X)
	intY := between(a.Y, b.Y, c.Y) || between(a.Y, b.Y, d.Y)

	return intX && intY
}

func between(from, to, x float64) bool {
	if from > to {
		from, to = to, from
	}
	return from <= x && x <= to
}

func sortPointsX(a, b Point) (Point, Point) {
	if a.X > b.X {
		return b, a
	}
	return a, b
}

func sign(x float64) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
