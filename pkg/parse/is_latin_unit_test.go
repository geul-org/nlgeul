//ff:func feature=parse type=helper control=sequence
//ff:what TestIsLatinUnit — a listed unit written in Latin letters (ASCII or fullwidth) is one, another text is not
package parse

import "testing"

func TestIsLatinUnit(t *testing.T) {
	if !isLatinUnit("kg") || isLatinUnit("py") || isLatinUnit("") {
		t.Fatal("isLatinUnit")
	}
	if !isLatinUnit("ｍ") || !isLatinUnit("ｇ") || !isLatinUnit("ｔ") || isLatinUnit("ｍｍ") || isLatinUnit("％") {
		t.Fatal("isLatinUnit fullwidth")
	}
}
