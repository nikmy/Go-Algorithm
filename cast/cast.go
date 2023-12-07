package cast

import (
	"golang.org/x/exp/constraints"

	"a.yandex-team.ru/library/go/slices"
)

type number interface {
	constraints.Integer | constraints.Float
}

func StaticNumber[From, To number](from From) To {
	return (To)(from)
}

func StaticNumbers[From, To number, S ~[]From](slice S) []To {
	return slices.Map(slice, StaticNumber[From, To])
}

type stringLike interface {
	~string | ~[]byte | ~[]rune
}

func StaticString[From stringLike](from From) string {
	return (string)(from)
}

func StaticStrings[From stringLike, S ~[]From](slice S) []string {
	return slices.Map(slice, StaticString[From])
}

func Dynamic[From any, To any](f From) To {
	if to, ok := any(f).(To); ok {
		return to
	}
	var empty To
	return empty
}

func DynamicSlice[FromElem any, ToElem any, S ~[]FromElem](slice S) []ToElem {
	if slice == nil {
		return nil
	}
	out := make([]ToElem, 0, len(slice))
	for _, elem := range slice {
		out = append(out, Dynamic[FromElem, ToElem](elem))
	}
	return out
}

func DynamicMap[
	FromKey comparable, FromValue any,
	ToKey comparable, ToValue any,
	M ~map[FromKey]FromValue,
](m M) map[ToKey]ToValue {
	if m == nil {
		return nil
	}
	out := make(map[ToKey]ToValue, len(m))
	for k, v := range m {
		out[Dynamic[FromKey, ToKey](k)] = Dynamic[FromValue, ToValue](v)
	}
	return out
}

func DynamicFunction[
	FromInput, FromOutput,
	ToInput, ToOutput any,
	F ~func(FromInput) FromOutput,
](f F) func(ToInput) ToOutput {
	if f == nil {
		return nil
	}
	return func(input ToInput) ToOutput {
		return Dynamic[FromOutput, ToOutput](f(Dynamic[ToInput, FromInput](input)))
	}
}
