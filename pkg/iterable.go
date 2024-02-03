package pkg

import (
	"go/ast"
	"reflect"

	"golang.org/x/tools/go/ast/astutil"
)



func detectIterables(ts *ast.TypeSpec) {
	var iterables = []*ast.StructType{}
	astutil.Apply(ts, nil, func(c *astutil.Cursor) bool {
		if c.Node() != nil {
			if st, ok := c.Node().(*ast.StructType); ok {
				var t ast.Expr
				for _, f := range st.Fields.List {
					if t == nil {
						t = f.Type
					} else if !compare(t, f.Type) {
						t = nil
						break
					}
				}
				
				if t != nil {
					iterables = append(iterables, st.Fields.List)
				}
				return false
			}
			
			reflect.Indirect(reflect.ValueOf(c.Parent())).FieldByName(c.Name())
			c.Replace(e.Name)
		}
		return true
	})
}
