package bucket

import "go.ufukty.com/gonfique/internal/tree"

type Bucket struct {
	heading string
	list    []string
	subs    []*Bucket
}

func New(heading string) *Bucket {
	return &Bucket{
		heading: heading,
		list:    []string{},
		subs:    []*Bucket{},
	}
}

func (b Bucket) count() int {
	c := len(b.list)
	for _, sub := range b.subs {
		c += sub.count()
	}
	return c
}

func (b Bucket) IsEmpty() bool {
	return b.count() == 0
}

func (b Bucket) String() string {
	if b.IsEmpty() {
		return ""
	}
	for _, i := range b.subs {
		if !i.IsEmpty() {
			b.list = append(b.list, i.String())
		}
	}
	return tree.List(b.heading, b.list)
}

func (b *Bucket) Add(s string) {
	b.list = append(b.list, s)
}

func (b *Bucket) Sub(heading string) *Bucket {
	sub := New(heading)
	b.subs = append(b.subs, sub)
	return sub
}
