//ff:func feature=token type=reader control=sequence
//ff:what TestFitsAny — article 2: every boundary between words fits, whatever the words
package token

import "testing"

func TestFitsAny(t *testing.T) {
	if !fitsAny(gap{}) || !fitsAny(gap{left: spacingWord("씨/열쇠:체언"), right: spacingWord("가/열쇠:동사"), leftLen: 1, rightLen: 1}) {
		t.Fatal("article 2")
	}
}
