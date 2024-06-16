package directives

import (
	"fmt"
	"go/ast"
	"reflect"
)

func (d *Directives) populateExprs() error {
	for kp, n := range d.Holders {
		switch n := n.(type) {
		case *ast.Field:
			d.TypeExprs[kp] = n.Type
		case *ast.ArrayType:
			d.TypeExprs[kp] = n.Elt
		default:
			return fmt.Errorf("unrecognized holder type: %s", reflect.TypeOf(n).String())
		}
	}
	return nil
}
