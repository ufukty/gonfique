package models

type Encoding string

var (
	Json = Encoding("json")
	Yaml = Encoding("yaml")
)

type Keypath string

type TypeName = string
