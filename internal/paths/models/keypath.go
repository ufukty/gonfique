package models

import "strings"

type FlattenKeypath string

func (kp FlattenKeypath) Terms() []string {
	if kp == "" {
		return []string{}
	}
	return strings.Split(string(kp), ".")
}

func (kp FlattenKeypath) Termination() string {
	terms := kp.Terms()
	if len(terms) == 0 {
		return ""
	}
	return terms[len(terms)-1]
}

func (kp FlattenKeypath) Parent() FlattenKeypath {
	terms := kp.Terms()
	if len(terms) <= 1 {
		return ""
	}
	return FlattenKeypath(strings.Join(terms[:len(terms)-1], "."))
}

func (kp FlattenKeypath) Sub(term string) FlattenKeypath {
	if kp == "" {
		return FlattenKeypath(term)
	}
	return FlattenKeypath(string(kp) + "." + term)
}
