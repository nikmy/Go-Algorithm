package geom

type Polygon []Point

func (pg Polygon) Size() int {
	return len(pg)
}

func (pg Polygon) Encloses(p Point) bool {
	r := Ray{p, p.VecTo(Point{p.X + 1, p.Y})}
	inside := false
	for i := 1; i < pg.Size(); i++ {
		down, top := pg[i-1], pg[i]
		switch {
		case down.Y == top.Y:
			continue
		case down.Y > top.Y:
			down, top = top, down
		}

		if r.Has(down) {
			continue
		}

		if r.Has(top) {
			inside = !inside
		}

		edge := Segment{A: down, B: top}
		if edge.IntersectRay(r) {
			inside = !inside
		}
	}
	return inside
}
