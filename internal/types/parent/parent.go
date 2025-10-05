package parent

import (
	"fmt"
	"go/ast"
	"go/token"
	"slices"

	"go.ufukty.com/gonfique/internal/datas/collects"
	"go.ufukty.com/gonfique/internal/files/config"
	"go.ufukty.com/gonfique/internal/namings"
	"go.ufukty.com/gonfique/internal/paths/models"
	"go.ufukty.com/gonfique/internal/transform"
)

type parentRefDetails struct {
	Fieldname transform.FieldName
}

func detailsForParentRefs(d *Directives) (map[config.Typename]parentRefDetails, error) {
	details := map[config.Typename]parentRefDetails{}

	for tn, kps := range d.instances {
		values := collects.WithSources[transform.FieldName, models.FlattenKeypath]{}
		for _, kp := range kps {
			dirs := d.directives[kp]
			if dirs.Parent != "" {
				values.Collect(dirs.Parent, kp)
			}
		}
		fieldname, err := values.One()
		if err == collects.ErrNoValues {
			continue
		} else if err != nil {
			return nil, fmt.Errorf("checking every keypath resolves to the typename %q and specifies a parent ref field name: %w", tn, err)
		}
		details[tn] = parentRefDetails{Fieldname: fieldname}
	}

	return details, nil
}

func selectExprForKeypath(d *Directives, kp models.FlattenKeypath) ast.Expr {
	var x ast.Expr = ast.NewIdent("c") // c is also hardcoded in coder.addReaderFunction
	ancestry := []models.FlattenKeypath{}
	for ckp := kp; ckp != ""; ckp = ckp.Parent() {
		ancestry = append(ancestry, ckp)
	}
	slices.Reverse(ancestry)
	for _, ancestor := range ancestry {
		x = &ast.SelectorExpr{
			X:   x,
			Sel: d.b.Fieldnames[d.holders[ancestor]].Ident(),
		}
	}
	return x
}

// Add parent fields to type expressions and generate value assignments
// for later embedding in ReadConfig function
func implementParentRefs() error {
	details, err := detailsForParentRefs(d)
	if err != nil {
		return fmt.Errorf("detailsForParentRefs: %w", err)
	}

	for tn, details := range details {
		pf := &ast.Field{
			Names: []*ast.Ident{details.Fieldname.Ident()},
			Type:  ast.NewIdent("any"),
			Tag:   &ast.BasicLit{Kind: token.STRING, Value: "`yaml:\"-\"`"},
		}
		ty := d.molds[tn].(*ast.StructType)
		ty.Fields.List = append(ty.Fields.List, pf)
	}

	sorted := slices.Clone(d.features.Parent)
	slices.Sort(sorted)
	d.b.ParentRefAssignStmts = []ast.Stmt{}
	for _, kp := range sorted {
		d.b.ParentRefAssignStmts = append(d.b.ParentRefAssignStmts, &ast.AssignStmt{
			Lhs: []ast.Expr{
				&ast.SelectorExpr{X: selectExprForKeypath(d, kp), Sel: ast.NewIdent("Parent")},
			},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.UnaryExpr{
				Op: token.AND,
				X:  selectExprForKeypath(d, kp.Parent()),
			}},
		})
	}

	for tn, details := range details {
		recvname := namings.Initial(string(tn))
		fd := &ast.FuncDecl{
			Recv: &ast.FieldList{List: []*ast.Field{{
				Names: []*ast.Ident{{Name: recvname}},
				Type:  tn.Ident(),
			}}},
			Name: ast.NewIdent("Get" + string(details.Fieldname)),
			Type: &ast.FuncType{
				Params:  &ast.FieldList{},
				Results: &ast.FieldList{List: []*ast.Field{{Type: ast.NewIdent("any")}}},
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{
				Results: []ast.Expr{&ast.SelectorExpr{
					X:   ast.NewIdent(recvname),
					Sel: details.Fieldname.Ident(),
				}},
			}}},
		}
		d.b.Accessors = append(d.b.Accessors, fd)
	}

	return nil
}
