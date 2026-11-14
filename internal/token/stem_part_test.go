//ff:func feature=token type=helper control=sequence
//ff:what TestStemPart — the first part of a word when it is a stem; the empty part for a word with no part or whose first part is a marker
package token

import "testing"

func TestStemPart(t *testing.T) {
	if got := stemPart(spacingWord("씨/열쇠:체언", "가/주격:가")); got.Text != "씨" || !got.Stem {
		t.Fatal(got)
	}
	if got := stemPart(spacingWord("-가/주격:가")); got.Text != "" || got.Stem {
		t.Fatal("a word of a marker alone", got)
	}
	if got := stemPart(Eojeol{}); got.Text != "" {
		t.Fatal("no part", got)
	}
}
