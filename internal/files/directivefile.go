package files

import (
	"fmt"
	"os"
	"slices"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
	"gopkg.in/yaml.v3"
)

type Directives struct {
	Named     models.TypeName    `yaml:"named"`
	Type      models.TypeName    `yaml:"type"`      // type-assigning
	Import    string             `yaml:"import"`    // type-assigning (optional)
	Embed     models.TypeName    `yaml:"embed"`     // type-defining
	Parent    models.FieldName   `yaml:"parent"`    // type-defining
	Accessors []models.FieldName `yaml:"accessors"` // type-defining
}

type DirectiveFile map[models.Keypath]Directives

func (df DirectiveFile) validate() error {
	for kp, dir := range df {
		typeAssigning := dir.Type != ""
		typeDefining := dir.Embed != "" && dir.Parent != "" && len(dir.Accessors) > 0
		if typeAssigning && typeDefining {
			return fmt.Errorf("the directives for %q includes both type defining and type assigning features", kp)
		}
	}
	return nil
}

func ReadDirectiveFile(path string) (*DirectiveFile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening the directive file to read: %w", err)
	}
	defer f.Close()
	df := &DirectiveFile{}
	err = yaml.NewDecoder(f).Decode(df)
	if err != nil {
		return nil, fmt.Errorf("decoding the directive file: %w", err)
	}
	if err := df.validate(); err != nil {
		return nil, fmt.Errorf("validating the directive file: %w", err)
	}
	return df, nil
}

func (df DirectiveFile) GetAccessors() map[models.Keypath][]models.FieldName {
	accessors := map[models.Keypath][]models.FieldName{}
	for kp, dirs := range df {
		if len(dirs.Accessors) > 0 {
			accessors[kp] = dirs.Accessors
		}
	}
	return accessors
}

func (df DirectiveFile) neededTypesForAccessorsDirective() []models.Keypath {
	needed := []models.Keypath{}
	for kp, drs := range df {
		if drs.Accessors != nil {
			needed = append(needed, kp) // struct
			for _, field := range drs.Accessors {
				needed = append(needed, kp.WithField(field)) // its field
			}
		}
	}
	return needed
}

func (df DirectiveFile) neededTypesForParentDirective() []models.Keypath {
	needed := []models.Keypath{}
	for kp, drs := range df {
		if drs.Parent != "" {
			needed = append(needed, kp.Parent())
		}
	}
	return needed
}

// both the struct and field types at each directive needs to be declared as named (not inline)
func (df DirectiveFile) NeededTypes() []models.Keypath {
	return datas.Uniq(slices.Concat(
		df.neededTypesForAccessorsDirective(),
		df.neededTypesForParentDirective(),
	))
}
