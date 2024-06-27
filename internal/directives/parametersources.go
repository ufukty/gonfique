package directives

import (
	"slices"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/directives/directivefile"
	"github.com/ufukty/gonfique/internal/models"
)

// use when parameters are already comparable (unlike slices, such as in accessors)
type comprParameterSources[K comparable] map[K][]models.WildcardKeypath

// this is for convenience opposed to sliceParameterDetailsSources.AddSource()
func (m comprParameterSources[K]) AddSource(parameters K, source models.WildcardKeypath) {
	if _, ok := m[parameters]; !ok {
		m[parameters] = []models.WildcardKeypath{}
	}
	m[parameters] = append(m[parameters], source)
}

// expensive map with comparable slice keys and values, use for accessors
type sliceTypeParameterSources[K comparable] map[*[]K][]models.WildcardKeypath

func (m sliceTypeParameterSources[K]) AddSource(parameters []K, source models.WildcardKeypath) {
	for s := range m {
		if slices.Equal(*s, parameters) {
			m[s] = append(m[s], source)
			return
		}
	}
	m[&parameters] = []models.WildcardKeypath{source}
}

type parameterSources struct {
	Accessors sliceTypeParameterSources[models.FieldPath]
	Declare   comprParameterSources[models.TypeName]
	Export    comprParameterSources[bool]
	Parent    comprParameterSources[directivefile.Parent]
	Replace   comprParameterSources[directivefile.Replace]
}

func (d *Directives) parameterSourceClassification() {
	d.ParameterSources = map[models.FlattenKeypath]parameterSources{}
	for kp, wckps := range datas.RevSliceMap(d.Expansions) {
		sources := parameterSources{
			Accessors: sliceTypeParameterSources[models.FieldPath]{},
			Declare:   comprParameterSources[models.TypeName]{},
			Export:    comprParameterSources[bool]{},
			Parent:    comprParameterSources[directivefile.Parent]{},
			Replace:   comprParameterSources[directivefile.Replace]{},
		}
		for _, wckp := range wckps {
			wckpdirectives := (*d.b.Df)[wckp]
			if len(wckpdirectives.Accessors) > 0 {
				sources.Accessors.AddSource(wckpdirectives.Accessors, wckp)
			}
			if wckpdirectives.Declare != "" {
				sources.Declare.AddSource(wckpdirectives.Declare, wckp)
			}
			sources.Export.AddSource(wckpdirectives.Export, wckp)
			if wckpdirectives.Parent.Fieldname != "" {
				sources.Parent.AddSource(wckpdirectives.Parent, wckp)
			}
			if wckpdirectives.Replace.Typename != "" {
				sources.Replace.AddSource(wckpdirectives.Replace, wckp)
			}
		}
		d.ParameterSources[kp] = sources
	}
}
