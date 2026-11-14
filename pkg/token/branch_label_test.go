//ff:func feature=token type=helper control=sequence
//ff:what TestBranchLabel — a branch is called by the labels of every part of its words in order, joined with " + "; a branch with no words by nothing
package token

import "testing"

func TestBranchLabel(t *testing.T) {
	b := Branch{Words: []Eojeol{
		{Parts: []Part{
			{Text: "고양이", Stem: true, Name: Name{"열쇠", "체언"}},
			{Text: "가", Name: Name{"주격", "가"}},
		}, End: 4},
		{Parts: []Part{{Text: "마신", Stem: true}}, End: 6},
	}}
	if got := b.Label(); got != "열쇠:체언 + 주격:가 + 새 낱말" {
		t.Fatalf("got %q", got)
	}
	if got := (Branch{}).Label(); got != "" {
		t.Fatalf("no words: %q", got)
	}
}
