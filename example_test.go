package tfconv_test

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nitrado/tfconv"
)

func Example() {
	type TestObject struct {
		Test string `json:"test"`
	}

	testSchema := map[string]*schema.Schema{
		"test": {
			Type: schema.TypeString,
		},
	}

	conv := tfconv.New("json")

	// This would be used in the Resource read functions.
	in := TestObject{
		Test: "Hello, World!",
	}
	data, err := conv.Flatten(in, testSchema)
	if err != nil {
		panic(err)
	}

	// This would be used in the Resource write functions.
	var out TestObject
	if err = conv.Expand(data, &out); err != nil {
		panic(err)
	}

	fmt.Printf("%v", out)

	// Output: {Hello, World!}
}
