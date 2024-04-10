package config

type Endpoint struct {
	Method, Path string
}

type Service struct {
	Path      string
	Endpoints []Endpoint
}
