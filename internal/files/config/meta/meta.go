package meta

type Meta struct {
	Type    string `yaml:"type"`
	Package string `yaml:"package"`
}

var Default = Meta{
	Package: "config",
	Type:    "Config",
}
