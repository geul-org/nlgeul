//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestUnitReaderNext — the stretches made first come first, then those of the net as it is worked out (the first before the net is done), then those of the signs after the letters, and then none
package token

import "testing"

func TestUnitReaderNext(t *testing.T) {
	r := streamReader(t, "words-far.tsv")
	u := unitsOf(r.ix, "…"+windowText(4)+".", false)[0]
	rd := newUnitReader(r.ix, u, false, &netStats{})
	first, _ := rd.next()
	second, _ := rd.next()
	if first.cells != nil || second.cells == nil || rd.run == nil || rd.run.j >= rd.run.n {
		t.Fatal("the net is worked out as it is asked for")
	}
	got := 2
	for _, ok := rd.next(); ok; _, ok = rd.next() {
		got++
	}
	last := len(rd.queue)
	if rd.run != nil || got < 5 || last != 0 {
		t.Fatal(got, rd.run)
	}
	if _, ok := rd.next(); ok {
		t.Fatal("none")
	}
}
