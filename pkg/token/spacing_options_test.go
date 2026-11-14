//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSpacingOptions — a reader is not made for a spacing rule that is not one of SpacingArticles, with or without a word table, and the error names it; every name of SpacingArticles is taken
package token

import (
	"strings"
	"testing"
)

func TestSpacingOptions(t *testing.T) {
	for _, words := range []string{"", "testdata/words-spacing.tsv"} {
		if _, err := New(Options{Words: words, Spacing: []string{"99"}}); err == nil || !strings.Contains(err.Error(), "99") {
			t.Fatal("a rule that is not one", words, err)
		}
		if _, err := New(Options{Words: words, Spacing: SpacingArticles}); err != nil {
			t.Fatal(err)
		}
	}
}
