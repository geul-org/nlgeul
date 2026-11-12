//ff:func feature=token type=helper control=sequence
//ff:what TestBranchString — a branch is printed as its words with a space between them; a branch with no words as nothing
package token

import "testing"

func TestBranchString(t *testing.T) {
	b := Branch{Words: []Eojeol{
		{Parts: []Part{
			{Text: "고양이", Stem: true, Name: Name{"열쇠", "체언"}},
			{Text: "가", Name: Name{"주격", "가"}},
		}, End: 4},
		{Parts: []Part{{Text: "마신", Stem: true}}, End: 6},
		{Parts: []Part{{Text: "물", Stem: true}, {Text: "을", Name: Name{"목적격", "을"}}}, End: 8},
	}}
	if got := b.String(); got != "고양이[열쇠:체언]+가[주격:가] 마신 물+을[목적격:을]" {
		t.Fatalf("got %q", got)
	}
	if got := (Branch{}).String(); got != "" {
		t.Fatalf("no words: %q", got)
	}
}
