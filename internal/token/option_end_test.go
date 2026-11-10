//ff:func feature=token type=helper control=sequence
//ff:what TestOptionEnd — one or two dashes and a Latin letter begin an option, which runs over letters, digits and - _ =; no dash, three, a digit or Hangul after the dashes, and dashes alone give 0
package token

import "testing"

func TestOptionEnd(t *testing.T) {
	if optionEnd([]rune("--dry-run=1을")) != 11 || optionEnd([]rune("-v")) != 2 {
		t.Fatal(optionEnd([]rune("--dry-run=1을")))
	}
	if optionEnd([]rune("v")) != 0 || optionEnd([]rune("---v")) != 0 || optionEnd([]rune("-3")) != 0 || optionEnd([]rune("-가")) != 0 || optionEnd([]rune("--")) != 0 {
		t.Fatal("no option")
	}
}
