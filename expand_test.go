package tfconv_test

import (
	"testing"

	terra "github.com/nitrado/tfconv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/resource"
)

func newInt(i int) *int {
	return &i
}

func TestConverter_Expand(t *testing.T) {
	c := terra.New("json")
	c.Register(resource.Quantity{}, func(v any) (any, error) {
		return resource.ParseQuantity(v.(string))
	}, func(v any) (any, error) {
		q := v.(resource.Quantity)
		return (&q).String(), nil
	})

	data := []any{map[string]any{
		"str":   "test-str",
		"alias": "test-alias",
		"int":   1,
		"float": 2.3,
		"bool":  true,
		"slice": []any{map[string]any{
			"a": "test-t",
		}, map[string]any{
			"a": "test-t-also",
		}},
		"map": map[string]any{
			"foo": 4,
		},
		"map_convert": map[string]any{
			"foo": "bar",
		},
		"struct": []any{map[string]any{
			"a": "test-ptr-t",
			"b": []any{map[string]any{
				"value": 16,
			}},
		}},
		"q": "205m",
		"q_ptr": []any{map[string]any{
			"value": "2Mi",
		}},
	}}

	got := TestObject{}
	err := c.Expand(data, &got)

	require.NoError(t, err)
	want := TestObject{
		Str:   "test-str",
		Alias: StrAlias("test-alias"),
		Int:   1,
		Float: 2.3,
		Bool:  true,
		Slice: []T{
			{A: "test-t"},
			{A: "test-t-also"},
		},
		Map: map[string]int{
			"foo": 4,
		},
		MapConvert: map[likeAString]string{
			"foo": "bar",
		},
		Struct: &T{
			A: "test-ptr-t",
			B: newInt(16),
			C: nil,
		},
		Q:    resource.MustParse("205m"),
		QPtr: ptrOf(resource.MustParse("2Mi")),
	}
	assert.Equal(t, want, got)
}
