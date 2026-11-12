//ff:func feature=token type=helper control=sequence
//ff:what TestPartString — a part prints what the letters do not show in brackets, its letters, and its candidates in square brackets; a stem no table knows prints bare, a marker with no candidates with empty brackets, a stem written wrong has the form it is written in after its candidates, and a part with no letters as what is written out alone
package token

import "testing"

func TestPartString(t *testing.T) {
	if got := (Part{Text: "다", Hidden: "ㄴ", Name: Name{"맺음", "ㄴ다"}}).String(); got != "(ㄴ)다[맺음:ㄴ다]" {
		t.Fatalf("hidden: %q", got)
	}
	if got := (Part{Text: "마신", Stem: true}).String(); got != "마신" {
		t.Fatalf("new word: %q", got)
	}
	if got := (Part{Text: "3", Stem: true, Name: Name{"수", ""}}).String(); got != "3[수]" {
		t.Fatalf("number: %q", got)
	}
	if got := (Part{Text: "x"}).String(); got != "x[]" {
		t.Fatalf("marker with no candidates: %q", got)
	}
	if got := (Part{Hidden: "이", Name: Name{"서술격", "이(빠짐)"}, Unseen: true}).String(); got != "(이)[서술격:이(빠짐)]" {
		t.Fatalf("written out only: %q", got)
	}
	if got := (Part{Text: "맡기", Written: "맞기", Stem: true, Name: Name{"열쇠", "동사"}}).String(); got != "맡기[열쇠:동사,오타정정:맞기]" {
		t.Fatalf("written wrong: %q", got)
	}
	if got := (Part{Text: "가", Name: Name{"주격", "가"}}).String(); got != "가[주격:가]" {
		t.Fatalf("marker: %q", got)
	}
}
