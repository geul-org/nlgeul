//ff:func feature=token type=helper control=sequence
//ff:what TestLetterLine — the letters from Start up to End; a place outside the sentence is kept inside it
package token

import "testing"

func TestLetterLine(t *testing.T) {
	l := []rune("가나다라")
	if got := letterLine(Span{1, 3}, l); got != "나다" {
		t.Fatal(got)
	}
	if got := letterLine(Span{2, 9}, l); got != "다라" {
		t.Fatal(got)
	}
	if got := letterLine(Span{-3, 1}, l); got != "가" {
		t.Fatal(got)
	}
	if got := letterLine(Span{7, 9}, l); got != "" {
		t.Fatal(got)
	}
	if got := letterLine(Span{3, 1}, l); got != "" {
		t.Fatal(got)
	}
}
