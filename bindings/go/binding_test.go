package tree_sitter_lyra_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_lyra "github.com/hassak-47/tree-sitter-lyra/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_lyra.Language())
	if language == nil {
		t.Errorf("Error loading Lyra grammar")
	}
}
