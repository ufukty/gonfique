package collects

import "iter"

type Set[E comparable] map[E]any

func (s *Set[E]) Add(e E) {
	(*s)[e] = new(any)
}

func (s *Set[E]) Remove(e E) {
	delete(*s, e)
}

func (s Set[E]) Iter() iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s {
			if !yield(v) {
				break
			}
		}
	}
}
