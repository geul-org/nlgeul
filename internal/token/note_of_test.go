//ff:func feature=token type=helper control=sequence
//ff:what TestNoteOf — a verdict above 0 is 통과; any other is 기각 with the reasons of the exceptions that applied, joined, and nothing after the colon when none applied
package token

import "testing"

func TestNoteOf(t *testing.T) {
	if got := noteOf(1, []string{"x"}); got != "통과" {
		t.Fatal(got)
	}
	if got := noteOf(-1, []string{"가", "나"}); got != "기각: 가, 나" {
		t.Fatal(got)
	}
	if got := noteOf(0, nil); got != "기각: " {
		t.Fatal(got)
	}
}
