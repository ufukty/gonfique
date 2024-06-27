package directives

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/directives/directivefile"
	"github.com/ufukty/gonfique/internal/matcher"
	"github.com/ufukty/gonfique/internal/models"
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

func (d *Directives) populateDirectivesForKeypaths() {
	l := map[models.FlattenKeypath]directivefile.Directives{}
	for wckp, dirs := range *d.b.Df {
		for _, kp := range d.Expansions[wckp] {
			d := l[kp]
			if len(dirs.Accessors) > 0 {
				d.Accessors = dirs.Accessors
			}
			if dirs.Declare != "" {
				d.Declare = dirs.Declare
			}
			if dirs.Embed.Typename != "" {
				d.Embed = dirs.Embed
			}
			d.Export = dirs.Export
			if dirs.Parent.Fieldname != "" {
				d.Parent = dirs.Parent
			}
			if dirs.Replace.Typename != "" {
				d.Replace = dirs.Replace
			}
			l[kp] = d
		}
	}
	d.DirectivesForKeypaths = l
}

type featuresForKeypaths struct {
	Accessors []models.FlattenKeypath
	Embed     []models.FlattenKeypath
	Export    []models.FlattenKeypath
	Import    []models.FlattenKeypath
	Declare   []models.FlattenKeypath
	Parent    []models.FlattenKeypath
	Replace   []models.FlattenKeypath
}

func (d *Directives) populateFeaturesForKeypaths() {
	for kp, dirs := range d.DirectivesForKeypaths {
		if len(dirs.Accessors) > 0 {
			d.FeaturesForKeypaths.Accessors = append(d.FeaturesForKeypaths.Accessors, kp)
		}
		if dirs.Declare != "" {
			d.FeaturesForKeypaths.Declare = append(d.FeaturesForKeypaths.Declare, kp)
		}
		if dirs.Embed.Typename != "" {
			d.FeaturesForKeypaths.Embed = append(d.FeaturesForKeypaths.Embed, kp)
		}
		if dirs.Export {
			d.FeaturesForKeypaths.Export = append(d.FeaturesForKeypaths.Export, kp)
		}
		if dirs.Parent.Fieldname != "" {
			d.FeaturesForKeypaths.Parent = append(d.FeaturesForKeypaths.Parent, kp)
		}
		if dirs.Replace.Typename != "" {
			d.FeaturesForKeypaths.Replace = append(d.FeaturesForKeypaths.Replace, kp)
		}
	}
}
