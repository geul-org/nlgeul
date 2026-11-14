//ff:func feature=token type=reader control=sequence
//ff:what TestMerged — candidates that end in the same word are one state that holds each state before them once, in the order they came, the first standing for them; the state of a word met before stays as it is when it comes again (its preds not changed); the root of any object is the scan's root; each state comes once in the order its word first came; the same word printed the same of other values is another state and a twin
package token

import "testing"

func TestMerged(t *testing.T) {
	sc := &scan{}
	root := sc.registry().root
	x := addState(sc, []*path{root}, wordStep(1, "x", 1, -1))
	y := addState(sc, []*path{root}, wordStep(1, "y", 1, -1))
	w := wordStep(3, "w", 1, -2)
	attachedRoot := &path{verdict: 1}
	out := sc.merged([]*path{grow(x, w), grow(y, w), grow(x, w), grow(attachedRoot, wordStep(3, "v", 1, -1)), grow(root, wordStep(3, "v", 1, -1)), grow(y, wordStep(3, "u", 1, -1))})
	if len(out) != 3 || out[0].last.sig != "w" || out[1].last.sig != "v" || out[2].last.sig != "u" {
		t.Fatal(out)
	}
	if len(out[0].preds) != 2 || out[0].preds[0] != x || out[0].preds[1] != y || out[0].prev != x {
		t.Fatal("preds", out[0].preds)
	}
	if len(out[1].preds) != 1 || out[1].preds[0] != root {
		t.Fatal("the root is the scan's root", out[1].preds)
	}
	again := sc.merged([]*path{out[0], out[0]})
	if len(again) != 1 || again[0] != out[0] || len(out[0].preds) != 2 {
		t.Fatal("a state met before", again, out[0].preds)
	}
	if sc.registry().twins != 0 {
		t.Fatal("no twins yet")
	}
	other := sc.merged([]*path{grow(x, wordStep(3, "w", 0.5, -2))})
	if sc.registry().twins != 1 || other[0] == out[0] {
		t.Fatal("a twin is another state", sc.registry().twins)
	}
}
