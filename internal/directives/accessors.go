package directives

import (
	"fmt"
	"go/ast"
	"go/token"
	"slices"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/models"
	"github.com/ufukty/gonfique/internal/namings"
)

func generateGetter(typename models.TypeName, fieldname models.FieldName, fieldtype models.TypeName) *ast.FuncDecl {
	recvname := namings.Initial(string(typename))
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{{
				Names: []*ast.Ident{{Name: recvname}},
				Type:  typename.Ident(),
			}},
		},
		Name: &ast.Ident{Name: "Get" + string(fieldname)},
		Type: &ast.FuncType{
			Params: &ast.FieldList{},
			Results: &ast.FieldList{
				List: []*ast.Field{{
					Type: fieldtype.Ident(),
				}},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.SelectorExpr{
							X:   ast.NewIdent(recvname),
							Sel: fieldname.Ident(),
						},
					},
				},
			},
		},
	}
}

func generateSetter(typename models.TypeName, fieldname models.FieldName, fieldtype models.TypeName) *ast.FuncDecl {
	recvname := namings.Initial(string(typename))
	paramname := "v"
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{{
				Names: []*ast.Ident{{Name: recvname}},
				Type:  &ast.StarExpr{X: typename.Ident()},
			}},
		},
		Name: &ast.Ident{Name: "Set" + string(fieldname)},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{{
					Names: []*ast.Ident{{Name: paramname}},
					Type:  fieldtype.Ident(),
				}},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.SelectorExpr{
						X:   ast.NewIdent(recvname),
						Sel: fieldname.Ident(),
					}},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{&ast.Ident{Name: paramname}},
				},
			},
		},
	}
}

func lookFieldTypename(b *bundle.Bundle, kp models.FlattenKeypath) (models.TypeName, error) {
	if slices.Contains(b.NeedsToBeNamed, kp) {
		tn, ok := b.GeneratedTypenames[kp]
		if !ok {
			return "", fmt.Errorf("generated typename is not found for keyapth: %s", kp)
		}
		return tn, nil
	} else if ident, ok := b.TypeExprs[kp].(*ast.Ident); ok {
		return models.TypeName(ident.Name), nil
	} else {
		return "", fmt.Errorf("type name is not found")
	}
}

func ImplementAccessors(b *bundle.Bundle) error {
	if b.Df == nil {
		return fmt.Errorf("directive file is not populated")
	} else if b.GeneratedTypenames == nil {
		return fmt.Errorf("typenames is missing")
	}
	b.Accessors = []*ast.FuncDecl{}

	for wildcardkp, directives := range *b.Df {
		if directives.Accessors != nil {
			matches, ok := b.Expansions[wildcardkp]
			if !ok {
				return fmt.Errorf("no match for keypath: %s", wildcardkp)
			}
			for _, match := range matches {
				kp, ok := b.Keypaths[match]
				if !ok {
					return fmt.Errorf("no flatten keypath found for wildcard keypath and match: %s / %s", wildcardkp, b.Keypaths[match])
				}

				structtypename, ok := b.GeneratedTypenames[kp]
				if !ok {
					return fmt.Errorf("can't find the assigned type name for struct: %s", wildcardkp)
				}

				for _, fieldpath := range directives.Accessors {

					fieldtypename, err := lookFieldTypename(b, kp.WithFieldPath(fieldpath))
					if err != nil {
						return fmt.Errorf("looking for correct typename: %w", err)
					}

					field := b.Holders[kp.WithFieldPath(fieldpath)]
					assignedFieldName := b.Fieldnames[field]

					b.Accessors = append(b.Accessors,
						generateGetter(structtypename, assignedFieldName, fieldtypename),
						generateSetter(structtypename, assignedFieldName, fieldtypename),
					)
				}
			}
		}
	}
	return nil
}
