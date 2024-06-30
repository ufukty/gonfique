package files

import (
	"fmt"
	"os"

	"github.com/ufukty/gonfique/internal/models"
	"gopkg.in/yaml.v3"
)

func ReadMappings(src string) (map[models.WildcardKeypath]models.TypeName, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	ms := map[models.WildcardKeypath]models.TypeName{}
	err = yaml.NewDecoder(f).Decode(&ms)
	if err != nil {
		return nil, fmt.Errorf("decoding: %w", err)
	}

	return ms, nil
}
