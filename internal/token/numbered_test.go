//ff:func feature=token type=helper control=sequence
//ff:what TestNumbered — each word gets its number in the order it first comes, a word that came before (whatever the words before it) keeps its number, and the next number goes on across lists
package token

import "testing"

func TestNumbered(t *testing.T) {
	driveAll(t)
	sc := &scan{}
	root := &path{verdict: 1}
	x, y := grow(root, wordStep(1, "x", 1, -1)), grow(root, wordStep(1, "y", 1, -1))
	a, b := grow(x, wordStep(3, "a", 1, -1)), grow(y, wordStep(3, "a", 1, -1))
	c := grow(x, wordStep(3, "c", 1, -1))
	sc.numbered([]*path{a, c, b})
	if a.last.born != 0 || c.last.born != 1 || b.last.born != 0 || sc.registry().next != 2 {
		t.Fatal(a.last.born, c.last.born, b.last.born, sc.registry().next)
	}
	d := grow(x, wordStep(2, "a", 1, -1))
	sc.numbered([]*path{d, a})
	if d.last.born != 2 || a.last.born != 0 || sc.registry().next != 3 {
		t.Fatal("a later list", d.last.born, a.last.born)
	}
}
