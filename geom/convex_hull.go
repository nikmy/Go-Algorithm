package geom

import "sort"

func ConvexHull(points []Point) []Point {
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p.X < q.X || p.X == q.X && p.Y < q.Y
	})

	var lower []Point
	for _, p := range points {
		for len(lower) >= 2 {
			a, b := lower[len(lower)-2], lower[len(lower)-1]
			if Vector.Cross(p.VecTo(a), p.VecTo(b)) > 0 {
				break
			}
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	var upper []Point
	for i := len(points) - 1; i >= 0; i-- {
		p := points[i]
		for len(upper) >= 2 {
			a, b := upper[len(upper)-2], upper[len(upper)-1]
			if Vector.Cross(p.VecTo(a), p.VecTo(b)) > 0 {
				break
			}
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}

	return append(lower[:len(lower)-1], upper[:len(upper)-1]...)
}
