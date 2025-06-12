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
			"a.b.c":     {Iterator: true},
		},
	}
	fmt.Println(File(f))
	// Output:
	// file
	// ╰─ rules
	//    ├─ <A>
	//    │  ╰─ directives
	//    │     ├─ checking 'dict' value: invalid value: "hello world"
	//    │     ╰─ type targeting rules can't contain directives for values
	//    │        ╰─ dict
	//    ├─ <A>.*.<B>
	//    │  ╰─ path
	//    │     ╰─ type segment after start: "<B>"
	//    ├─ <A>.<B>
	//    │  ├─ path
	//    │  │  ╰─ type segment after start: "<B>"
	//    │  ╰─ directives
	//    │     ╰─ directives are missing
	//    ╰─ a.b.c
	//       ╰─ directives
	//          ╰─ value targeting rules can't contain directives for types
	//             ╰─ iterator
}
