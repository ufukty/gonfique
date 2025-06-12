package namings

import "strings"

func Initial(name string) string {
	return strings.ToLower(string(([]rune(name))[0]))
}
