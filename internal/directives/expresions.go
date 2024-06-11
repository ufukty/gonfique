package directives

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/ufukty/gonfique/internal/bundle"
)

func populateExprs(b *bundle.Bundle) error {
	for kp, n := range b.Holders {
		switch n := n.(type) {
		case *ast.Field:
			b.TypeExprs[kp] = n.Type
		case *ast.ArrayType:
			b.TypeExprs[kp] = n.Elt
		default:
			return fmt.Errorf("unrecognized holder type: %s", reflect.TypeOf(n).String())
		}
	}
	return nil
}
