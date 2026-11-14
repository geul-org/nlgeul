//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestOntoReader — a sign of a pair with letters written onto it is read as the contests of the sign alone, the stretches of the letters after it (worked out as they are asked for, with the words of the whole unit planted), and the signs after the letters; a head with no letters after it has no net, and the trace lines are those of the sign's reading
package token

import "testing"

func TestOntoReader(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	onto := unit{text: ")를.", head: 1, name: pairName, candidates: []pairing{{shutName, 1}}}
	rd := ontoReader(r.ix, onto, false, &netStats{})
	if rd.run == nil || rd.run.plant == nil || rd.run.first || len(rd.queue) == 0 {
		t.Fatal("the letters have a net beside the unit as a whole")
	}
	var cells int
	for st, ok := rd.next(); ok; st, ok = rd.next() {
		if st.cells != nil {
			cells++
		}
	}
	if cells == 0 {
		t.Fatal("no stretch of the letters")
	}
	bare := ontoReader(r.ix, unit{text: ")", head: 1, name: pairName, candidates: []pairing{{shutName, 1}}}, true, &netStats{})
	if bare.run != nil || len(bare.queue) == 0 {
		t.Fatal("no letters after the sign", bare.run)
	}
}
