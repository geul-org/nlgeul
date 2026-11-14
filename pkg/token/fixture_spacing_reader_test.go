//ff:func feature=token type=helper control=sequence
//ff:what spacingReader — a reader over a word table of testdata (none when the name is empty) with the given spacing rules on
package token

import "testing"

func spacingReader(t *testing.T, words string, spacing ...string) *Reader {
	t.Helper()
	if words != "" {
		words = "testdata/" + words
	}
	r, err := New(Options{Words: words, Spacing: spacing})
	if err != nil {
		t.Fatal(err)
	}
	return r
}
