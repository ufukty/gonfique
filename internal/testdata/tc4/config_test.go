package config

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func find(root ast.Node, ident ast.Ident) bool {
	found := false
	ast.Inspect(root, func(n ast.Node) bool {
		if n != nil {
			if n, ok := n.(*ast.Ident); ok {
				if n.Name == ident.Name {
					found = true
				}
			}
		}
		return !found
	})
	return found
}

func TestConfig(t *testing.T) {
	cfg, err := ReadConfig("config.yml")
	if err != nil {
		t.Fatal(fmt.Errorf("reading config: %w", err))
	}

	if cfg.Github.Domain != "github.com" {
		t.Fatal(fmt.Errorf("expected %q got %q", "github.com", cfg.Github.Domain))
	}

	f, err := parser.ParseFile(token.NewFileSet(), "config.go", nil, parser.AllErrors)
	if err != nil {
		t.Fatal(fmt.Errorf("parsing output file: %w", err))
	}

	if find(f, ast.Ident{Name: "Service"}) {
		t.Fatal("output file should include name of types that input file has no matching parts with its type")
	}
}
