package fieldlist

import "go/ast"

// Separates multiple field definitions per Field instance into isolated Field entries
func Flatten(fl *ast.FieldList) *ast.FieldList {
	flttn := &ast.FieldList{}
	for _, f := range fl.List {
		for _, ident := range f.Names {
			flttn.List = append(flttn.List, &ast.Field{
				Names: []*ast.Ident{ident},
				Type:  f.Type,
				Tag:   f.Tag,
			})
		}
	}
	return flttn
}
