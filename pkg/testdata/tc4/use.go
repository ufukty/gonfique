package config

type Endpoint struct {
	Method string
	Path   string
}

type Service struct {
	Path      string
	Endpoints []Endpoint
}
