package directivefile

import (
	"fmt"
	"os"

	"github.com/ufukty/gonfique/internal/models"
	"gopkg.in/yaml.v3"
)

type Embed struct {
	Typename   models.TypeName `yaml:"typename"`
	ImportPath string          `yaml:"import-path"`
	ImportAs   string          `yaml:"import-as"`
}

type Parent struct {
	Fieldname models.FieldName `yaml:"fieldname"`
	Accessors bool             `yaml:"accessors"`
	Level     int              `yaml:"level"`
}

type Replace struct {
	Typename   models.TypeName `yaml:"typename"`
	ImportPath string          `yaml:"import-path"`
	ImportAs   string          `yaml:"import-as"`
}

type Directives struct {
	Accessors []models.FieldPath `yaml:"accessors"`
	Embed     Embed              `yaml:"embed"`
	Export    bool               `yaml:"export"`
	Declare   models.TypeName    `yaml:"declare"`
	Parent    Parent             `yaml:"parent"`
	Replace   Replace            `yaml:"replace"`
}

type DirectiveFile map[models.WildcardKeypath]Directives

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
	return df, nil
}
