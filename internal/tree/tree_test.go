package tree

import (
	"fmt"
)

func ExampleList_single() {
	lines := []string{"Lorem", "Ipsum", "Dolor", "Sit", "Amet"}
	fmt.Println(List("Placeholders", lines))
	// Output:
	// Placeholders
	// ├─ Lorem
	// ├─ Ipsum
	// ├─ Dolor
	// ├─ Sit
	// ╰─ Amet
}

func ExampleList_multiple() {
	m := List("Placeholders", []string{"Lorem", "Ipsum", "Dolor", "Sit", "Amet"})
	m = List("Sets of placeholders", []string{m, m, m})
	fmt.Println(m)
	// Output:
	// Sets of placeholders
	// ├─ Placeholders
	// │  ├─ Lorem
	// │  ├─ Ipsum
	// │  ├─ Dolor
	// │  ├─ Sit
	// │  ╰─ Amet
	// ├─ Placeholders
	// │  ├─ Lorem
	// │  ├─ Ipsum
	// │  ├─ Dolor
	// │  ├─ Sit
	// │  ╰─ Amet
	// ╰─ Placeholders
	//    ├─ Lorem
	//    ├─ Ipsum
	//    ├─ Dolor
	//    ├─ Sit
	//    ╰─ Amet
}
