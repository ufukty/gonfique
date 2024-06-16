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

func (d *Directives) TypenameRequirementsForAccessors() error {
	for wckp, drs := range *d.b.Df {
		if drs.Accessors != nil {
			kps, ok := d.Expansions[wckp]
			if !ok {
				return fmt.Errorf("expansion is found for: %s", wckp)
			}
			for _, kp := range kps {
				d.NeededToBeReferred = append(d.NeededToBeReferred, kp) // struct
				for _, field := range drs.Accessors {
					d.NeededToBeReferred = append(d.NeededToBeReferred, kp.WithFieldPath(field)) // its field
				}
			}
		}
	}
	return nil
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

func (d *Directives) AddAccessorFuncDecls() error {
	if d.b.Df == nil {
		return fmt.Errorf("directive file is not populated")
	} else if d.ElectedTypenames == nil {
		return fmt.Errorf("elected type names are missing")
	}
	d.b.Accessors = []*ast.FuncDecl{}

	fieldsfortypes := map[models.TypeName]map[models.FieldName]models.TypeName{}
	for wckp, directives := range *d.b.Df {
		if directives.Accessors != nil {
			for _, kp := range d.Expansions[wckp] {
				tn := d.ElectedTypenames[kp]
				if _, ok := fieldsfortypes[tn]; !ok {
					fieldsfortypes[tn] = map[models.FieldName]models.TypeName{}
				}
				for _, fp := range directives.Accessors {
					fkp := kp.WithFieldPath(fp)
					ftn := d.ElectedTypenames[fkp]
					fn := d.b.Fieldnames[d.Holders[fkp]]
					fieldsfortypes[tn][fn] = ftn
				}
			}
		}
	}

	sorted := maps.Keys(fieldsfortypes)
	slices.SortFunc(sorted, caseInsensitiveCompareTypenames)
	for _, tn := range sorted {
		fields := fieldsfortypes[tn]
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
