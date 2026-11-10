//ff:func feature=mini type=reader control=sequence
//ff:what TestReadWord — 이것은 has a chain and a tail; 봤어요. is read without its full stop and through 요; 매우 has neither; 요 alone is not a tail
package mini

import "testing"

func TestReadWord(t *testing.T) {
	sent := &sentence{}
	w := readWord("이것은", false, sent)
	if w.chain.size != 2 || len(w.tails) != 1 || w.tails[0].kind != kindAdn || w.longest != 2 || w.last || w.sent != sent {
		t.Fatalf("%+v", w)
	}
	w = readWord("봤어요.", true, sent)
	if w.text != "봤어요" || w.longest != 5 || !w.last {
		t.Fatalf("%+v", w)
	}
	if w = readWord("매우", false, sent); w.chain.size != 0 || w.tails != nil {
		t.Fatalf("%+v", w)
	}
	if w = readWord("요", false, sent); w.tails != nil {
		t.Fatalf("%+v", w)
	}
}
