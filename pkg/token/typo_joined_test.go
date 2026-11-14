//ff:func feature=token type=reader control=sequence
//ff:what TestTypoJoined — the first line of a written form is a misspelt form of its own; a line with the same right form joins its kind to it; a line with another right form is added after it
package token

import "testing"

func TestTypoJoined(t *testing.T) {
	got := typoJoined(nil, "맡기", kindVerb)
	if len(got) != 1 || got[0] != (typo{"맡기", kindVerb}) {
		t.Fatal(got)
	}
	got = typoJoined(got, "맡기", kindNominal)
	if len(got) != 1 || got[0].kinds != kindVerb|kindNominal {
		t.Fatal(got)
	}
	got = typoJoined(got, "맞이", kindNominal)
	if len(got) != 2 || got[0].right != "맡기" || got[1] != (typo{"맞이", kindNominal}) {
		t.Fatal(got)
	}
}
