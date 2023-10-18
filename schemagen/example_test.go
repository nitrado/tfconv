package schemagen_test

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nitrado/tfconv/schemagen"
)

func Example() {
	type TestObject struct {
		Test string `json:"test"`
	}

	docs := func(obj any, sf reflect.StructField) string {
		// Return the documentation for the struct field.
		return ""
	}

	customize := func(obj any, sf *reflect.StructField, typ reflect.Type, s *schema.Schema) (string, reflect.Kind, bool) {
		// Customize the schema output for a sctruct field.
		return "", typ.Kind(), true
	}

	gen := schemagen.New(docs, customize, "json")

	s, err := gen.Struct(TestObject{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", s)

	// Output: map[test:{
	// Type: schema.TypeString,
	// }]
}
