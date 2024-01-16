package dsu

type Forest struct {
	root []int
	rank []int
	Size int
}

func NewSetForest(nComponents int) *Forest {
	f := &Forest{make([]int, nComponents), make([]int, nComponents), 0}
	for i := 0; i < nComponents; i++ {
		f.root[i] = i
		f.rank[i] = 1
	}
	f.Size = nComponents
	return f
}

func (f *Forest) Find(x int) int {
	if x != f.root[x] {
		f.root[x] = f.Find(f.root[x])
	}
	return f.root[x]
}

func (f *Forest) SameSet(x, y int) bool {
	return f.Find(x) == f.Find(y)
}

func (f *Forest) Union(x, y int) {
	x, y = f.Find(x), f.Find(y)
	if f.rank[y] < f.rank[x] {
		f.root[y] = x
	} else {
		f.root[x] = y
	}
	if f.rank[y] == f.rank[x] {
		f.rank[y]++
	}
	f.Size--
}
