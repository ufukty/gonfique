package directives

import (
	"github.com/ufukty/gonfique/internal/directives/directivefile"
	"github.com/ufukty/gonfique/internal/models"
)

func (d *Directives) linearizeDirectives() {
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
	d.Directives = l
}
