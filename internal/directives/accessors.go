package directives

import (
	"fmt"
	"go/ast"
	"go/token"
	"slices"

	"github.com/ufukty/gonfique/internal/models"
	"github.com/ufukty/gonfique/internal/namings"
	"golang.org/x/exp/maps"
)

type accessordetails map[models.FieldName]models.TypeName

func accessorDetailsForTypes(d *Directives) (map[models.TypeName]accessordetails, error) {
	details := map[models.TypeName]accessordetails{}

	for tn, kps := range d.Instances {
		init := true
		details[tn] = map[models.FieldName]models.TypeName{}
		for _, kp := range kps {
			for _, fn := range d.Directives[kp].Accessors {
				fkp := kp.WithFieldPath(fn)
				ftn := d.Elected[fkp]
				fn := d.b.Fieldnames[d.Holders[fkp]]
				current := details[tn]
				if !init {
					if current[fn] != ftn {
						return nil, fmt.Errorf("typename %q is directed to have accessors on the field %q  which its type resolving to different types", tn, fn)
					}
				}
				details[tn][fn] = ftn
				init = true
			}
		}
	}

	return details, nil
}

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
	if recvname == "v" {
		paramname = "value"
	}
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

func (d *Directives) addAccessorFuncDecls() error {
	details, err := accessorDetailsForTypes(d)
	if err != nil {
		return fmt.Errorf("accessorDetailsForTypes: %w", err)
	}
	tns := maps.Keys(details)
	slices.SortFunc(tns, caseInsensitiveCompareTypenames)

	d.b.Accessors = []*ast.FuncDecl{}
	for _, tn := range tns {
		fields := details[tn]
		sortedfields := maps.Keys(fields)
		slices.SortFunc(sortedfields, caseInsensitiveCompareFieldnames)
		for _, fn := range sortedfields {
			ftn := fields[fn]
			d.b.Accessors = append(d.b.Accessors,
				generateGetter(tn, fn, ftn),
				generateSetter(tn, fn, ftn),
			)
		}
	}

	return nil
}
