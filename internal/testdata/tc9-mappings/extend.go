package config

import "fmt"

func (e Endpoint) String() string { // if this doesn't throw error in compilation, test will complete
	return fmt.Sprintf("%s %s", e.Method, e.Path)
}
