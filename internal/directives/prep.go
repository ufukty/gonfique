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
	d.Keypaths = resolver.keypaths
	d.Holders = datas.Invmap(d.Keypaths)
}

func (d *Directives) populateExprs() error {
	for kp, n := range d.Holders {
		switch n := n.(type) {
		case *ast.Field:
			d.KeypathTypeExprs[kp] = n.Type
		case *ast.ArrayType:
			d.KeypathTypeExprs[kp] = n.Elt
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
			kps = append(kps, d.Keypaths[match])
		}
		d.Expansions[kp] = kps
	}
	return nil
}

func (d *Directives) populateDirectivesAndFeaturesForKeypaths() {
	dfk := map[models.FlattenKeypath]directivefile.Directives{}
	for kp, sources := range d.ParameterSources {
		ds := directivefile.Directives{}
		if len(sources.Accessors) > 0 {
			ds.Accessors = *maps.Keys(sources.Accessors)[0]
			d.FeaturesForKeypaths.Accessors = append(d.FeaturesForKeypaths.Accessors, kp)
		}
		if len(sources.Declare) > 0 {
			ds.Declare = maps.Keys(sources.Declare)[0]
			d.FeaturesForKeypaths.Declare = append(d.FeaturesForKeypaths.Declare, kp)
		}
		if len(sources.Embed) > 0 {
			ds.Embed = maps.Keys(sources.Embed)[0]
			d.FeaturesForKeypaths.Embed = append(d.FeaturesForKeypaths.Embed, kp)
		}
		if len(sources.Export) > 0 {
			ds.Export = maps.Keys(sources.Export)[0]
			d.FeaturesForKeypaths.Export = append(d.FeaturesForKeypaths.Export, kp)
		}
		if len(sources.Parent) > 0 {
			ds.Parent = maps.Keys(sources.Parent)[0]
			d.FeaturesForKeypaths.Parent = append(d.FeaturesForKeypaths.Parent, kp)
		}
		if len(sources.Replace) > 0 {
			ds.Replace = maps.Keys(sources.Replace)[0]
			d.FeaturesForKeypaths.Replace = append(d.FeaturesForKeypaths.Replace, kp)
		}
		dfk[kp] = ds
	}
	d.DirectivesForKeypaths = dfk
}
