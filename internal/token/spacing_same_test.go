//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSpacingSame — a reader with the rules on gives the same contests, word ends and marks each time it reads the same input
package token

import (
	"slices"
	"testing"
)

func TestSpacingSame(t *testing.T) {
	for _, in := range []string{"갑돌씨가", "뚜루뚜겸총무다", "삼억오천만원을", "이백삼십개", "솔·겸총무", "아버지가방에들어가신다"} {
		a, b := spacingReader(t, "words-spacing.tsv", SpacingArticles...), spacingReader(t, "words-spacing.tsv", SpacingArticles...)
		if !slices.Equal(a.Lines(in, true), b.Lines(in, true)) || !slices.Equal(a.Cuts(in), b.Cuts(in)) || !slices.Equal(a.Marks(in), b.Marks(in)) || !slices.Equal(a.Marks(in), a.Marks(in)) {
			t.Fatal(in)
		}
	}
}
