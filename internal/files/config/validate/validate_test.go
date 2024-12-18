package validate

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/files/config/meta"
)

func ExampleValidate() {
	f := &config.File{
		Meta: meta.Default,
		Paths: map[config.Path]config.PathConfig{
			"<A>":       {Dict: "hello world"},
			"<A>.<B>":   {Dict: config.Struct},
			"<A>.*.<B>": {Dict: config.Struct},
		},
		Types: map[config.Typename]config.TypeConfig{},
	}
	fmt.Println(File(f))
	// Output:
	// file
	// ╰─ path section
	//    ├─ <A>
	//    │  ╰─ checking `dict` directive value: invalid value
	//    │
	//    ├─ <A>.<B>
	//    │  ╰─ paths should not contain a typename after the beginning: <B>
	//    │
	//    ╰─ <A>.*.<B>
	//       ╰─ paths should not contain a typename after the beginning: <B>
}
