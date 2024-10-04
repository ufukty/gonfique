package directives

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/directives/directivefile"
	"github.com/ufukty/gonfique/internal/matcher"
	"github.com/ufukty/gonfique/internal/models"
	"golang.org/x/exp/maps"
)

func (d *Directives) populateKeypathsAndHolders() {
	resolver := newresolver(d.b.OriginalKeys)
	resolver.dfs(d.b.CfgType, nil, []string{})
	d.keypaths = resolver.keypaths
	d.holders = datas.Invmap(d.keypaths)
}

func (d *Directives) populateExprs() error {
	for kp, n := range d.holders {
		switch n := n.(type) {
		case *ast.Field:
			d.exprs[kp] = n.Type
		case *ast.ArrayType:
			d.exprs[kp] = n.Elt
		default:
			return fmt.Errorf("unrecognized holder type: %T", n)
		}
	}
	return nil
}

func (d *Directives) expandKeypaths() error {
	for kp := range *d.b.Df {
		matches, err := matcher.FindTypeDefHoldersForKeypath(d.b.CfgType, kp, d.b.OriginalKeys)
		if err != nil {
			return fmt.Errorf("matching the rule: %w", err)
		}
		if len(matches) == 0 {
			fmt.Printf("No match for keypath: %s\n", kp)
		}
		kps := []models.FlattenKeypath{}
		for _, match := range matches {
			kps = append(kps, d.keypaths[match])
		}
		d.expansions[kp] = kps
	}
	return nil
}

func (d *Directives) populateDirectivesAndFeaturesForKeypaths() {
	dfk := map[models.FlattenKeypath]directivefile.Directives{}
	for kp, sources := range d.sources {
		ds := directivefile.Directives{}
		if len(sources.Accessors) > 0 {
			ds.Accessors = *maps.Keys(sources.Accessors)[0]
			d.features.Accessors = append(d.features.Accessors, kp)
		}
		if len(sources.Declare) > 0 {
			ds.Declare = maps.Keys(sources.Declare)[0]
			d.features.Declare = append(d.features.Declare, kp)
		}
		if len(sources.Embed) > 0 {
			ds.Embed = maps.Keys(sources.Embed)[0]
			d.features.Embed = append(d.features.Embed, kp)
		}
		if len(sources.Export) > 0 {
			ds.Export = maps.Keys(sources.Export)[0]
			d.features.Export = append(d.features.Export, kp)
		}
		if len(sources.Parent) > 0 {
			ds.Parent = maps.Keys(sources.Parent)[0]
			d.features.Parent = append(d.features.Parent, kp)
		}
		if len(sources.Replace) > 0 {
			ds.Replace = maps.Keys(sources.Replace)[0]
			d.features.Replace = append(d.features.Replace, kp)
		}
		dfk[kp] = ds
	}
	d.directives = dfk
}
