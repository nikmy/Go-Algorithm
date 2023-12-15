package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindInSeparated(t *testing.T) {
	type args struct {
		s          string
		substr     string
		delimiters string
	}

	type testcase struct {
		name string

		args args

		want bool
	}

	cases := [...]testcase{
		{
			name: "single value",
			args: args{
				s:          "test",
				substr:     "test",
				delimiters: ",",
			},
			want: true,
		},
		{
			name: "multiple values",
			args: args{
				s:          "test1,test2,test3",
				substr:     "test2",
				delimiters: ",",
			},
			want: true,
		},
		{
			name: "last value",
			args: args{
				s:          "test1,test2,test3",
				substr:     "test3",
				delimiters: ",",
			},
			want: true,
		},
		{
			name: "not in list",
			args: args{
				s:          "test1,test2,test3",
				substr:     "test42",
				delimiters: ",",
			},
			want: false,
		},
		{
			name: "not in list but prefix",
			args: args{
				s:          "test1,test2,test3",
				substr:     "test",
				delimiters: ",",
			},
			want: false,
		},
		{
			name: "not in list but middle",
			args: args{
				s:          "_test_,test2,test3",
				substr:     "test",
				delimiters: ",",
			},
			want: false,
		},
		{
			name: "not in list but suffix",
			args: args{
				s:          "tes_test,test2,test3",
				substr:     "test",
				delimiters: ",",
			},
			want: false,
		},
		{
			name: "not in list too long",
			args: args{
				s:          "test1,test2,test3",
				substr:     "very_very_long_string",
				delimiters: ",",
			},
			want: false,
		},
		{
			name: "multiple delimiters",
			args: args{
				s:          "test1,test2;test3,test4",
				substr:     "test3",
				delimiters: ",;",
			},
			want: true,
		},
		{
			name: "empty string",
			args: args{
				s:          "",
				substr:     "test",
				delimiters: ",",
			},
			want: false,
		},
		{
			name: "empty substr",
			args: args{
				s:          "test",
				substr:     "",
				delimiters: ",",
			},
			want: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, FindInSeparated(tt.args.s, tt.args.substr, tt.args.delimiters))
		})
	}
}
