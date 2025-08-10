package skiplist

import (
	"context"
	"runtime"
	"slices"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestList_Find(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name string
		list *List
		elem int64
		want bool
	}{
		{
			name: "found max",
			list: Make(generateInts(1, 9, 1)...),
			elem: 8,
			want: true,
		},
		{
			name: "found min",
			list: Make(generateInts(1, 9, 1)...),
			elem: 1,
			want: true,
		},
		{
			name: "found middle",
			list: Make(generateInts(1, 9, 1)...),
			elem: 4,
			want: true,
		},
		{
			name: "not found",
			list: Make(generateInts(1, 9, 1)...),
			elem: 9,
			want: false,
		},
		{
			name: "found negative",
			list: Make(1, -2, -5, 7, -8),
			elem: -5,
			want: true,
		},
		{
			name: "not found negative",
			list: Make(1, -2, -5, 7, -8),
			elem: -4,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("dump:\n%s\n", tt.list.SDump())
			if got := tt.list.Find(tt.elem); got != tt.want {
				t.Errorf("List.Find(%d) = %v, want %v", tt.elem, got, tt.want)
			}
		})
	}
}

func TestList_Insert(t *testing.T) {
	t.Parallel()

	l := Make(1, 2, 3)
	if !l.Insert(4) {
		t.Error("cannot insert new element: expected true, got false")
	}
	if l.Insert(3) {
		t.Error("double insertion: expected false, got true")
	}
}

func TestList_Delete(t *testing.T) {
	t.Parallel()

	l := Make(1, 2, 3)
	if l.Delete(4) {
		t.Errorf("deleted element that does not exist")
	}
	if !l.Delete(1) {
		t.Errorf("cannot delete element that exists")
	}
	if l.Delete(1) {
		t.Errorf("double deleted element")
	}
}

func TestList_Elements(t *testing.T) {
	t.Parallel()

	t.Run("empty list", func(t *testing.T) {
		list := New()
		for elem := range list.Elements {
			t.Errorf("Unexpected element %d", elem)
		}
	})

	t.Run("traversal", func(t *testing.T) {
		l := Make(generateInts(1, 11, 1)...)
		want := int64(1)
		for elem := range l.Elements {
			if elem != want {
				t.Errorf("List.Elements() = %v, want %v", elem, want)
			}
			want++
		}
	})
}

func TestList_Visual(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name string
		list *List
	}{
		{
			name: "empty list",
			list: New(),
		},
		{
			name: "one element",
			list: Make(42),
		},
		{
			name: "small",
			list: Make(generateInts(6, 0, -1)...),
		},
		{
			name: "medium",
			list: Make(generateInts(-1, 11, 1)...),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("fmt: %s\n", tt.list)
			t.Logf("dump:\n%s\n", tt.list.SDump())
		})
	}
}

func TestList_RaceFree(t *testing.T) {
	t.Run("find-find", func(t *testing.T) {
		ctrl := newCtrl(t)
		list := Make(generateInts(2, 11, 2)...)
		for i := int64(2); i <= 10; i += 2 {
			ctrl.Spawn(200, func() {
				assert.Truef(t, list.Find(i), "missing element %d", i)
			})
		}
		for i := int64(1); i <= 9; i += 2 {
			ctrl.Spawn(200, func() { assert.Falsef(t, list.Find(i), "unexpected element %d", i) })
		}

		ctrl.Run(5 * time.Second)
	})

	t.Run("insert-insert", func(t *testing.T) {
		ctrl := newCtrl(t)
		list := New()
		inserted := [50]atomic.Int64{}
		for i := int64(0); i < 50; i++ {
			ctrl.Spawn(400, func() {
				ok := list.Insert(i)
				if ok {
					n := int(inserted[i].Add(1))
					assert.Lessf(t, n, 2, "multiple insertion: %d", n)
				}
			})
		}

		ctrl.Run(5 * time.Second)

		assert.Equal(t, slices.Collect(list.Elements), generateInts(0, 50, 1))

		t.Logf("list after all operations:\n%s\n", list.SDump())
	})

	t.Run("find-insert", func(t *testing.T) {
		ctrl := newCtrl(t)
		list := New()
		inserted := [50]atomic.Int64{}
		for i := int64(0); i < 50; i++ {
			ctrl.Spawn(400, func() {
				_ = list.Find(i)

				ok := list.Insert(i)
				if ok {
					n := int(inserted[i].Add(1))
					assert.Lessf(t, n, 2, "multiple insertion: %d", n)
				}

				ok = list.Find(i)
				assert.Truef(t, ok, "can't found element %d after insertion", i)
			})
		}

		ctrl.Run(5 * time.Second)

		t.Logf("list after all operations:\n%s\n", list.SDump())
	})

	t.Run("delete-delete", func(t *testing.T) {
		ctrl := newCtrl(t)
		list := Make(generateInts(0, 60, 1)...)

		t.Logf("list before all operations:\n%s\n", list.SDump())

		for i := int64(0); i < 60; i++ {
			ctrl.Spawn(400, func() {
				_ = list.Delete(i * 2)
				ok := list.Delete(i * 2)
				if ok {
					assert.Failf(t, "double delete", "element %d deleted after delete", i*2)
				}

				_ = list.Delete(i*2 + 1)
				ok = list.Delete(i*2 + 1)
				if ok {
					assert.Failf(t, "double delete", "element %d deleted after delete", i*2+1)
				}
			})
		}

		ctrl.Run(5 * time.Second)

		t.Logf("list after all operations:\n%s\n", list.SDump())
	})

	t.Run("find-delete", func(t *testing.T) {
		ctrl := newCtrl(t)
		list := Make(generateInts(0, 40, 2)...)

		t.Logf("list before all operations:\n%s\n", list.SDump())

		deleted := [50]atomic.Int64{}
		for i := int64(0); i < 50; i++ {
			ctrl.Spawn(400, func() {
				_ = list.Find(i * 2)

				ok := list.Delete(i*2 + 1)
				assert.Falsef(t, ok, "deleted element %d that does not exist", i*2+1)

				ok = list.Delete(i * 2)
				if ok {
					n := int(deleted[i].Add(1))
					assert.Lessf(t, n, 2, "multiple delete of %d: %d", i*2, n)
				}

				if ok {
					assert.Falsef(t, list.Find(i*2), "element %d found after deletion", i*2)
				}
			})
		}

		ctrl.Run(5 * time.Second)

		t.Logf("list after all operations:\n%s\n", list.SDump())
	})

	t.Run("insert-delete", func(t *testing.T) {
		ctrl := newCtrl(t)
		list := New()

		t.Logf("list before all operations:\n%s\n", list.SDump())

		counts := [100][2]atomic.Int64{}
		inc := func(i int64, b bool) {
			if b {
				counts[i][0].Add(1)
			}
		}
		dec := func(i int64, b bool) {
			if b {
				counts[i][1].Add(1)
			}
		}

		for i := int64(0); i < 50; i++ {
			ctrl.Spawn(10, func() {
				inc(i*2+1, list.Insert(i*2+1))
				inc(i*2, list.Insert(i*2))
				dec(i*2, list.Delete(i*2))
				dec(i*2+1, list.Delete(i*2+1))
			})
		}

		ctrl.Run(5 * time.Second)

		for i := range &counts {
			inserts := counts[i][0].Load()
			deletes := counts[i][1].Load()
			assert.Equal(t, inserts, deletes, "inserts (expected) and deletes (actual) disbalanced")
			assert.Greaterf(t, inserts, int64(0), "no successful inserts / deletes")
		}

		assert.Len(t, slices.Collect(list.Elements), 0, "list must be empty")

		t.Logf("list after all operations:\n%s\n", list.SDump())
	})
}

func generateInts(start, end, step int) []int64 {
	ints := make([]int64, 0, (end-start)/step+1)
	for i := start; i < end; i += step {
		ints = append(ints, int64(i))
	}
	return ints
}

func newCtrl(t *testing.T) *pController {
	return &pController{
		t:  t,
		do: make(chan struct{}),
	}
}

type pController struct {
	t  *testing.T
	wg sync.WaitGroup
	do chan struct{}
	n  atomic.Int64
}

func (c *pController) Run(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	close(c.do)
	done := make(chan struct{})
	go func() {
		defer close(done)
		c.wg.Wait()
	}()
	select {
	case <-done:
	case <-ctx.Done():
		stack := make([]byte, 100_000)
		n := runtime.Stack(stack, true)
		assert.Failf(c.t, "timed out", "%d goroutines stuck:\n--- stacktrace ---\n%s", c.n.Load(), stack[:n])
	}
}

func (c *pController) Spawn(n int, g func()) {
	c.wg.Add(n)
	c.n.Add(int64(n))
	for i := 0; i < n; i++ {
		go func() {
			defer c.wg.Done()
			<-c.do
			g()
			c.n.Add(-1)
		}()
	}
}
