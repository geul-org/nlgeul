//ff:func feature=token type=helper control=sequence
//ff:what TestLastKind — the kind of the last part of a word; "" for a word with no part
package token

import "testing"

func TestLastKind(t *testing.T) {
	if got := lastKind(spacingWord("읽/열쇠:동사", "어/맺음:어")); got != "맺음" {
		t.Fatal(got)
	}
	if got := lastKind(Eojeol{}); got != "" {
		t.Fatal(got)
	}
}
