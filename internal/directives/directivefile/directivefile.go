package directivefile

import (
	"fmt"
	"os"

	"github.com/ufukty/gonfique/internal/models"
	"gopkg.in/yaml.v3"
)

type Directives struct {
	Named     models.TypeName    `yaml:"named"`
	Type      models.TypeName    `yaml:"type"`   // type-assigning
	Import    string             `yaml:"import"` // type-assigning (optional)
	Export    bool               `yaml:"export"`
	Embed     models.TypeName    `yaml:"embed"`     // type-defining
	Parent    models.FieldName   `yaml:"parent"`    // type-defining
	Accessors []models.FieldPath `yaml:"accessors"` // type-defining
}

type DirectiveFile map[models.WildcardKeypath]Directives

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
