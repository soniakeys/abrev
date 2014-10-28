package abrev_test

import (
	"fmt"

	"github.com/soniakeys/abrev"
)

func Example() {
	cmd := abrev.Set{"open", "export", "expand", "pan", "please"}
	fmt.Println(cmd.Expand("op")) // unique
	fmt.Println(cmd.Expand("o"))  // two words contain an "o"
	fmt.Println(cmd.Expand("xr")) // a subsequence is not a substring
	fmt.Println(cmd.Expand("xn"))
	fmt.Println(cmd.Expand("pan")) // a full match excludes partial matches
	fmt.Println(cmd.Expand("plz")) // no magic
	// Output:
	// [open]
	// [open export]
	// [export]
	// [expand]
	// [pan]
	// []
}
