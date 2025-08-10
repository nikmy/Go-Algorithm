package skiplist

import (
	"testing"
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
			list: Make(1, 2, 3, 4, 5, 6, 7, 8),
			elem: 8,
			want: true,
		},
		{
			name: "found min",
			list: Make(1, 2, 3, 4, 5, 6, 7, 8),
			elem: 1,
			want: true,
		},

		{
			name: "found middle",
			list: Make(1, 2, 3, 4, 5, 6, 7, 8),
			elem: 4,
			want: true,
		},
		{
			name: "not found",
			list: Make(1, 2, 3, 4, 5, 6, 7, 8),
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
				t.Errorf("List.Find() = %v, want %v", got, tt.want)
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
		l := Make(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
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
			list: Make(6, 5, 4, 3, 2, 1),
		},
		{
			name: "medium",
			list: Make(-1, -2, -3, -4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("fmt: %s\n", tt.list)
			t.Logf("dump:\n%s\n", tt.list.SDump())
		})
	}
}
