package tfconv

import (
	"reflect"
	"strings"

	"github.com/ettle/strcase"
)

// ConverterFn is a function that can convert a value.
type ConverterFn func(v any) (any, error)

type conversion struct {
	expand  ConverterFn
	flatten ConverterFn
}

// NameFunc is a function used to determine a field name.
type NameFunc = func(name string) string

// Converter converts Terraform formatted data to and from
// object types.
type Converter struct {
	tag string

	conversions map[reflect.Type]conversion
	nameFn      NameFunc
}

// New returns a new converter.
func New(tag string) *Converter {
	return NewWithName(nil, tag)
}

// NewWithName returns a new converter with the give nameFn.
func NewWithName(nameFn NameFunc, tag string) *Converter {
	if tag == "" {
		tag = "json"
	}
	if nameFn == nil {
		nameFn = strcase.ToSnake
	}

	return &Converter{
		tag:         tag,
		conversions: map[reflect.Type]conversion{},
		nameFn:      nameFn,
	}
}

// Register registers a custom type conversion.
func (c *Converter) Register(v any, expand, flatten ConverterFn) {
	t := reflect.TypeOf(v)
	c.conversions[t] = conversion{
		expand:  expand,
		flatten: flatten,
	}
}

func (c *Converter) resolveName(sf reflect.StructField) string {
	jsonName := sf.Tag.Get(c.tag)
	if name, _, ok := strings.Cut(jsonName, ","); ok {
		jsonName = name
	}
	if jsonName != "" && jsonName != "-" {
		return c.nameFn(jsonName)
	}
	return c.nameFn(sf.Name)
}
