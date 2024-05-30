package appendix

import (
	"go/ast"
	"reflect"
	"strings"
	"testing"
)

func TestKeys(t *testing.T) {
	for node, key := range Keys {
		if field, ok := node.(*ast.Field); ok {
			if st, ok := reflect.StructTag(strings.ReplaceAll(field.Tag.Value, "`", "")).Lookup("yaml"); !ok {
				t.Fatalf("can't get the value for struct tag 'yaml' in %q for key %q", field.Tag.Value, key)
			} else if st != key {
				t.Fatalf("expected %q got %q", st, key)
			}
		} else {
			t.Fatal("non-Field AST node")
		}
	}
}
