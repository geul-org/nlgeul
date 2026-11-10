//ff:func feature=token type=helper control=sequence
//ff:what TestHiddenOf — a marker written with brackets at its head parts into what the letters do not show and the letters; one with no brackets shows all its letters
package token

import "testing"

func TestHiddenOf(t *testing.T) {
	if hidden, shown := hiddenOf("(ㄴ)다"); hidden != "ㄴ" || shown != "다" {
		t.Fatalf("(ㄴ)다: %q %q", hidden, shown)
	}
	if hidden, shown := hiddenOf("(이)"); hidden != "이" || shown != "" {
		t.Fatalf("(이): %q %q", hidden, shown)
	}
	if hidden, shown := hiddenOf("가"); hidden != "" || shown != "가" {
		t.Fatalf("가: %q %q", hidden, shown)
	}
}
