//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReadUnitNet — a unit with signs at its edges is the sign contests of the old reading around the stretches the net read (the signs have no cells, the letters have); a sign of a pair with letters written onto it is the old reading as it is; the trace lines are the old reading's, once each
package token

import "testing"

func TestReadUnitNet(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	u := unitsOf(r.ix, "…아버지가방에.", false)[0]
	got, _ := readUnitNet(r.ix, u, false)
	if len(got) != 3 || got[0].cells != nil || got[1].cells == nil || got[2].cells != nil || got[0].contest.Text != "…" || got[2].contest.Text != "." {
		t.Fatal(got)
	}
	onto := unit{text: ")를.", head: 1, name: pairName, candidates: []pairing{{shutName, 1}}}
	old, _ := readUnit(r.ix, onto, false)
	got, _ = readUnitNet(r.ix, onto, false)
	if len(got) != len(old) || len(got) != 3 || got[1].cells != nil || got[1].contest.Text != old[1].Text || len(got[1].contest.Branches) != len(old[1].Branches) {
		t.Fatal("written onto", got)
	}
	traced := unit{text: "아버지가방에", notes: []string{"head"}}
	_, want := readUnit(r.ix, traced, true)
	_, notes := readUnitNet(r.ix, traced, true)
	if len(notes) == 0 || len(notes) != len(want) || notes[0] != "head" {
		t.Fatal("the trace", len(notes), len(want))
	}
	for i := range want {
		if notes[i] != want[i] {
			t.Fatal("trace line", i, notes[i], want[i])
		}
	}
}
