//ff:func feature=parse type=reader control=sequence
//ff:what TestFacts — every fact of the word is carried over, tails and sentence included; a word without a sentence gets an empty one
package parse

import "testing"

func TestFacts(t *testing.T) {
	w := readSentence("물이 얼음이 되었다")[1]
	w.nextBound = true
	f := facts(w)
	if f.Text != "얼음이" || f.Chain.Marked != "얼음+이" || f.Chain.Kase != w.chain.kase || f.Last || !f.NextBound || f.Sent.Nominative != 2 || f.Sent.Stems[0] != "되" {
		t.Fatalf("%+v", f)
	}
	f = facts(readWord("먹었다", true, nil))
	if len(f.Tails) != 1 || f.Tails[0].Marked != "먹+었+다" || f.Longest != 4 || !f.Last || f.Sent == nil {
		t.Fatalf("%+v", f)
	}
}
