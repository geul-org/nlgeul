//ff:func feature=token type=reader control=sequence
//ff:what TestSettle — the words are numbered; then a scan that does not build the net gives the candidates best first, the same print once (top), and one that builds it gives the states of the net (merged)
package token

import "testing"

func TestSettle(t *testing.T) {
	root := &path{verdict: 1}
	mk := func() []*path {
		return []*path{grow(root, wordStep(1, "a", 1, -3)), grow(root, wordStep(1, "b", 1, -1)), grow(root, wordStep(1, "a", 1, -3))}
	}
	sc := &scan{ix: scanOf("x").ix}
	got := sc.settle(mk())
	if len(got) != 2 || got[0].last.sig != "b" || got[1].last.sig != "a" || got[0].last.born != 1 || got[1].last.born != 0 {
		t.Fatal(got)
	}
	net := &scan{ix: sc.ix, netted: true}
	got = net.settle(mk())
	if len(got) != 2 || got[0].last.sig != "a" || got[1].last.sig != "b" || len(got[0].preds) != 1 || len(net.registry().all) != 2 {
		t.Fatal("the net", got)
	}
}
