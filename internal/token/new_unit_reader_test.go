//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNewUnitReader — a unit is read as it is asked for: the signs before the letters come first, then the stretches of the net, then the signs after, and the trace lines are those of readUnit (the unit's own first); a unit of letters only has no signs around it; a sign of a pair with letters written onto it is read in two
package token

import "testing"

func TestNewUnitReader(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	u := unitsOf(r.ix, "…아버지가방에.", false)[0]
	rd := newUnitReader(r.ix, u, true, &netStats{})
	var texts []string
	for st, ok := rd.next(); ok; st, ok = rd.next() {
		texts = append(texts, st.contest.Text)
	}
	if len(texts) != 3 || texts[0] != "…" || texts[2] != "." || len(rd.take()) == 0 {
		t.Fatal(texts)
	}
	plain := newUnitReader(r.ix, unitsOf(r.ix, "아버지가방에", false)[0], false, &netStats{})
	if st, ok := plain.next(); !ok || st.cells == nil || len(plain.queue) != 0 || plain.run == nil {
		t.Fatal("letters only")
	}
	onto := newUnitReader(r.ix, unit{text: ")를.", head: 1, name: pairName, candidates: []pairing{{shutName, 1}}}, false, &netStats{})
	if st, ok := onto.next(); !ok || st.cells != nil {
		t.Fatal("the sign comes first", st)
	}
}
