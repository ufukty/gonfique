package replace

type Agent struct {
	Imports []string
}

func New() *Agent {
	return &Agent{
		Imports: []string{},
	}
}
