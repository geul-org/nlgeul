//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunNext — the stretches of a long unit come one at a time and the first is cut off before the unit's end is reached; they cover the letters of the unit from the first to the last without a gap; a unit with no letters gives none
package token

import "testing"

func TestNetRunNext(t *testing.T) {
	r, _ := windowRun(t, "words-far.tsv", windowText(5))
	first, ok := r.next()
	if !ok || r.j >= r.n || first.contest.Text != "고양이가" {
		t.Fatal(r.j, r.n, first.contest.Text)
	}
	text := first.contest.Text
	for st, ok := r.next(); ok; st, ok = r.next() {
		text += st.contest.Text
	}
	if text != windowText(5) || r.made != len(r.stats.widths)+1 {
		t.Fatal(text, r.made)
	}
	if _, ok := r.next(); ok {
		t.Fatal("the end")
	}
	empty := newNetRun(&scan{ix: r.sc.ix}, unit{}, true, nil, &netStats{})
	if _, ok := empty.next(); ok {
		t.Fatal("no letters")
	}
}
