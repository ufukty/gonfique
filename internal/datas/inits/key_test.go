package inits

import (
	"testing"
)

func TestKey(t *testing.T) {
	m := map[string][]int{}
	Key(m, "")
	m[""] = append(m[""], 0)
}

func TestKey2(t *testing.T) {
	m := map[string]map[string][]int{}
	Key2(m, "", "")
	m[""][""] = append(m[""][""], 0)
}

func TestKey3(t *testing.T) {
	m := map[string]map[string]map[string][]int{}
	Key3(m, "", "", "")
	m[""][""][""] = append(m[""][""][""], 0)
}
