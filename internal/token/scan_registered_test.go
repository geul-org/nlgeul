//ff:func feature=token type=helper control=sequence
//ff:what TestScanRegistered — a candidate becomes a state: the state before it as its only pred and its prev, where its word begins, its print the word's alone, kept in the registry in order; a second state for the same word printed the same (of other values) counts as a twin
package token

import "testing"

func TestScanRegistered(t *testing.T) {
	sc := &scan{}
	x := sc.registered(grow(&path{verdict: 1}, wordStep(2, "x", 1, -1)), stateKey{word: wordKey{0, 2, "x"}})
	if len(x.preds) != 1 || x.preds[0] != sc.registry().root || x.prev != sc.registry().root || x.start != 0 || x.sig != "x" || len(sc.registry().all) != 1 || sc.registry().twins != 0 {
		t.Fatal(x)
	}
	y := sc.registered(grow(x, wordStep(5, "y", 1, -1)), stateKey{word: wordKey{2, 5, "y"}})
	if y.start != 2 || y.preds[0] != x || y.sig != "y" || len(sc.registry().all) != 2 || sc.registry().twins != 0 {
		t.Fatal(y)
	}
	z := sc.registered(grow(x, wordStep(5, "y", 0.5, -2)), stateKey{word: wordKey{2, 5, "y"}, verdict: 0.5, pref: -2})
	if sc.registry().twins != 1 || len(sc.registry().all) != 3 || sc.registry().states[stateKey{word: wordKey{2, 5, "y"}, verdict: 0.5, pref: -2}] != z {
		t.Fatal("a twin", sc.registry().twins)
	}
}
