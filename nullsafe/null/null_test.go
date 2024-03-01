package null

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFast(t *testing.T) {
	type testCase struct {
		runner func(*testing.T)
		name   string
	}

	cases := [...]testCase{
		{
			name: "fulfilled value",
			runner: func(t *testing.T) {
				v := Value[int](42)
				require.False(t, v.IsNull())
				require.NotPanics(t, func() { _ = v.Must() })
				require.Equal(t, 42, *v.Must())
			},
		},
		{
			name: "set to null",
			runner: func(t *testing.T) {
				late := Null[int]()
				late.Set(42)
				require.NotPanics(t, func() { _ = late.Must() })
				require.Equal(t, 42, *late.Must())
			},
		},
		{
			name: "panic on unsafe access",
			runner: func(t *testing.T) {
				require.Panics(t, func() { _ = Null[int]().Must() })
			},
		},
	}

	for _, tt := range &cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.runner(t)
		})
	}
}
