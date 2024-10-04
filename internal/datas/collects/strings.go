package collects

import "fmt"

func join[S ~string](ss []S, sep string) string {
	j := ""
	for _, s := range ss {
		j = fmt.Sprintf("%s%s%s", j, sep, s)
	}
	return j
}
