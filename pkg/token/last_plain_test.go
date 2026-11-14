//ff:func feature=token type=helper control=sequence
//ff:what TestLastPlain — the place after the last letter that is no Hangul syllable; 0 when all are
package token

import "testing"

func TestLastPlain(t *testing.T) {
	if lastPlain([]rune("a/b를")) != 3 || lastPlain([]rune("가나")) != 0 || lastPlain([]rune("가a")) != 2 || lastPlain(nil) != 0 {
		t.Fatal(lastPlain([]rune("a/b를")))
	}
}
