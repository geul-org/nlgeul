//ff:func feature=token type=helper control=sequence
//ff:what streamReader — a reader over the word table of testdata given by its file name (the small table for streams is words-stream.tsv; none when the name is empty)
package token

import "testing"

func streamReader(t *testing.T, words string) *Reader {
	t.Helper()
	if words != "" {
		words = "testdata/" + words
	}
	r, err := New(Options{Words: words})
	if err != nil {
		t.Fatal(err)
	}
	return r
}
