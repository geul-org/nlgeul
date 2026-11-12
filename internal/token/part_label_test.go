//ff:func feature=token type=helper control=sequence
//ff:what TestPartLabel — a stem no table knows is called 새 낱말; a known stem and a marker are called by their readings, a stem written wrong with the form it is written in, and signs with no name by their letters
package token

import "testing"

func TestPartLabel(t *testing.T) {
	if got := (Part{Text: "마신", Stem: true}).Label(); got != "새 낱말" {
		t.Fatalf("new word: %q", got)
	}
	if got := (Part{Text: "물", Stem: true, Names: []Name{{"열쇠", "체언"}, {"열쇠", "동사"}}}).Label(); got != "열쇠:체언|동사" {
		t.Fatalf("known stem: %q", got)
	}
	if got := (Part{Text: "맡기", Written: "맞기", Stem: true, Names: []Name{{"열쇠", "동사"}}}).Label(); got != "열쇠:동사,오타정정:맞기" {
		t.Fatalf("written wrong: %q", got)
	}
	if got := (Part{Text: "->", Stem: true}).Label(); got != "->" {
		t.Fatalf("signs with no name: %q", got)
	}
	if got := (Part{Text: "'"}).Label(); got != "" {
		t.Fatalf("a marker with no name: %q", got)
	}
	if got := (Part{Text: "(", Stem: true, Names: []Name{{"열기", "2"}}, Depth: 2}).Label(); got != "열기:2" {
		t.Fatalf("a sign of a pair: %q", got)
	}
	if got := (Part{Text: "가", Names: []Name{{"주격", "가"}, {"보격", "가"}}}).Label(); got != "주격:가|보격:가" {
		t.Fatalf("marker: %q", got)
	}
}
