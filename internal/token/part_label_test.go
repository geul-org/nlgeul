//ff:func feature=token type=helper control=sequence
//ff:what TestPartLabel — a stem no table knows is called 새 낱말; a known stem and a marker are called by their readings
package token

import "testing"

func TestPartLabel(t *testing.T) {
	if got := (Part{Text: "마신", Stem: true}).Label(); got != "새 낱말" {
		t.Fatalf("new word: %q", got)
	}
	if got := (Part{Text: "물", Stem: true, Names: []Name{{"열쇠", "체언"}, {"열쇠", "동사"}}}).Label(); got != "열쇠:체언|동사" {
		t.Fatalf("known stem: %q", got)
	}
	if got := (Part{Text: "가", Names: []Name{{"주격", "가"}, {"보격", "가"}}}).Label(); got != "주격:가|보격:가" {
		t.Fatalf("marker: %q", got)
	}
}
