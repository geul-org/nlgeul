//ff:func feature=parse type=helper control=sequence
//ff:what TestNameRune — Latin letters, digits and . ~ _ - are a name's letters; Hangul and other signs are not
package parse

import "testing"

func TestNameRune(t *testing.T) {
	if !nameRune('a') || !nameRune('7') || !nameRune('.') || !nameRune('~') || !nameRune('_') || !nameRune('-') || nameRune('가') || nameRune('/') {
		t.Fatal("name letters")
	}
}
