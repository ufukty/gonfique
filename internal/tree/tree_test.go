package tree

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
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

func TestListSecondLevel(t *testing.T) {
	lines := []string{"Lorem", "Ipsum", "Dolor", "Sit", "Amet"}
	m := List("Placeholders", lines)
	lines2 := []string{m, m, m}
	fmt.Println(List("Sets of placeholders", lines2))
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
