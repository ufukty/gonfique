// Code generated by govalid. DO NOT EDIT.

package config

import "fmt"

func (d Dict) Validate() error {
	switch d {
	case Map:
		return nil
	case Struct:
		return nil
	}
	return fmt.Errorf("invalid value: %q", d)
}