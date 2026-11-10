//ff:func feature=token type=reader control=sequence
//ff:what TestBareBases — ㅁ under three syllables, the other consonants under five
package token

import "testing"

func TestBareBases(t *testing.T) {
	if len(bareBases(jongM)) != 3 || len(bareBases(jongN)) != 5 {
		t.Fatal("bareBases")
	}
}
