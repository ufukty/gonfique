package validate

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/files/config/meta"
)

func ExampleFile() {
	f := &config.File{
		Meta: meta.Default,
		Rules: map[config.Path]config.Directives{
			"<A>":       {Dict: "hello world"},
			"<A>.<B>":   {},
			"<A>.*.<B>": {Export: true},
		},
	}
	fmt.Println(File(f))
	// Output:
	// file
	// ╰─ rules
	//    ├─ <A>
	//    │  ╰─ directives
	//    │     ╰─ checking 'dict' value: invalid value: "hello world"
	//    ├─ <A>.<B>
	//    │  ╰─ rule
	//    │     ╰─ type segment after start: <B>
	//    ╰─ <A>.*.<B>
	//       ╰─ rule
	//          ╰─ type segment after start: <B>
}
