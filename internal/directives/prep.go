package directives

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/directives/directivefile"
	"github.com/ufukty/gonfique/internal/matcher"
	"github.com/ufukty/gonfique/internal/models"
	"github.com/ufukty/gonfique/internal/namings"
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
			d.TypeExprs[kp] = n.Type
		case *ast.ArrayType:
			d.TypeExprs[kp] = n.Elt
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
			if dirs.Named != "" {
				d.Named = dirs.Named
			}
			if dirs.Parent != "" {
				d.Parent = dirs.Parent
			}
			if dirs.Embed != "" {
				d.Embed = dirs.Embed
			}
			if dirs.Type != "" {
				d.Type = dirs.Type
			}
			if dirs.Import != "" {
				d.Import = dirs.Import
			}
			if len(dirs.Accessors) > 0 {
				d.Accessors = dirs.Accessors
			}
			l[kp] = d
		}
	}
	d.DirectivesForKeypaths = l
}

func (d *Directives) populateTypenamesAutogen() {
	d.TypenamesAutogen = namings.GenerateTypenames(maps.Values(d.Keypaths))
}

func (d *Directives) populateTypenamesProvided() {
	d.TypenamesProvided = map[models.FlattenKeypath]models.TypeName{}
	for wckp, dirs := range *d.b.Df {
		if dirs.Named != "" {
			kps := d.Expansions[wckp]
			for _, kp := range kps {
				d.TypenamesProvided[kp] = dirs.Named
			}
		}
	}
}

type perfeature[T any] struct {
	Named     T
	Parent    T
	Type      T
	Import    T
	Embed     T
	Accessors T
}

func (d *Directives) populateFeaturesForKeypaths() {
	for kp, dirs := range d.DirectivesForKeypaths {
		if dirs.Named != "" {
			d.FeaturesForKeypaths.Named = append(d.FeaturesForKeypaths.Named, kp)
		}
		if dirs.Parent != "" {
			d.FeaturesForKeypaths.Parent = append(d.FeaturesForKeypaths.Parent, kp)
		}
		if dirs.Type != "" {
			d.FeaturesForKeypaths.Type = append(d.FeaturesForKeypaths.Type, kp)
		}
		if dirs.Import != "" {
			d.FeaturesForKeypaths.Import = append(d.FeaturesForKeypaths.Import, kp)
		}
		if dirs.Embed != "" {
			d.FeaturesForKeypaths.Embed = append(d.FeaturesForKeypaths.Embed, kp)
		}
		if len(dirs.Accessors) > 0 {
			d.FeaturesForKeypaths.Accessors = append(d.FeaturesForKeypaths.Accessors, kp)
		}
	}
}

func (d *Directives) populateFeaturesForTypenames() {
	for kp, dirs := range d.DirectivesForKeypaths {
		if dirs.Named != "" {
			d.FeaturesForKeypaths.Named = append(d.FeaturesForKeypaths.Named, kp)
		}
		if dirs.Parent != "" {
			d.FeaturesForKeypaths.Parent = append(d.FeaturesForKeypaths.Parent, kp)
		}
		if dirs.Type != "" {
			d.FeaturesForKeypaths.Type = append(d.FeaturesForKeypaths.Type, kp)
		}
		if dirs.Import != "" {
			d.FeaturesForKeypaths.Import = append(d.FeaturesForKeypaths.Import, kp)
		}
		if dirs.Embed != "" {
			d.FeaturesForKeypaths.Embed = append(d.FeaturesForKeypaths.Embed, kp)
		}
		if len(dirs.Accessors) > 0 {
			d.FeaturesForKeypaths.Accessors = append(d.FeaturesForKeypaths.Accessors, kp)
		}
	}
}
func (d *Directives) parentEnabledTypenames() []models.TypeName {

	enabled := map[models.TypeName]bool{}
	for wckp, dirs := range *d.b.Df {
		if dirs.Parent != "" {
			for _, kp := range d.Expansions[wckp] {
				enabled[d.TypenamesElected[kp]] = true
			}
		}
	}
	return maps.Keys(enabled)
}
