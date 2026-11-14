//ff:func feature=token type=helper control=iteration dimension=1
//ff:what netReaders — the readers the net is tried with: a small table of words written apart, a bigger one, the table of the spacing rules with them all on and all off, and no table; each with a name
package token

import "testing"

func netReaders(t *testing.T) map[string]*Reader {
	t.Helper()
	out := map[string]*Reader{}
	for name, o := range map[string]Options{
		"stream":      {Words: "testdata/words-stream.tsv"},
		"words":       {Words: "testdata/words.tsv"},
		"spacing-on":  {Words: "testdata/words-spacing.tsv", Spacing: SpacingArticles},
		"spacing-off": {Words: "testdata/words-spacing.tsv"},
		"words-on":    {Words: "testdata/words.tsv", Spacing: SpacingArticles},
		"none":        {},
	} {
		r, err := New(o)
		if err != nil {
			t.Fatal(name, err)
		}
		out[name] = r
	}
	return out
}
