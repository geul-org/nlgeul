//ff:func feature=token type=helper control=sequence
//ff:what TestIsLatinUnit — a listed unit is one, another text is not
package token

import "testing"

func TestIsLatinUnit(t *testing.T) {
	if !isLatinUnit("kg") || isLatinUnit("py") || isLatinUnit("") {
		t.Fatal("isLatinUnit")
	}
}
