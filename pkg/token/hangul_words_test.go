//ff:func feature=token type=helper control=sequence
//ff:what TestHangulWords — the words whose first letter is a Hangul syllable, in their order; a word of Latin letters and an empty word are left out
package token

import (
	"slices"
	"testing"
)

func TestHangulWords(t *testing.T) {
	got := hangulWords([]string{"개", "GB", "", "m", "마리", "%"})
	if !slices.Equal(got, []string{"개", "마리"}) {
		t.Fatal(got)
	}
	if got := hangulWords(nil); len(got) != 0 {
		t.Fatal(got)
	}
}
