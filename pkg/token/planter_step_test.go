//ff:func feature=token type=reader control=sequence
//ff:what TestPlanterStep — the places of a unit that is the sign of a pair alone are worked out and the state that is the sign alone is noted as one that has nothing left when the sign is taken off; for a unit with letters after the sign no such state is made; the units of the tests are read by the net too
package token

import "testing"

func TestPlanterStep(t *testing.T) {
	driveAll(t)
	r := streamReader(t, "words-far.tsv")
	alone := unit{text: ")", head: 1, name: pairName, candidates: []pairing{{shutName, 1}}}
	pl := newPlanter(unitScan(r.ix, alone), 1)
	if len(pl.transparent) != 1 || len(pl.best[1]) != 1 {
		t.Fatal("the sign alone", len(pl.transparent), len(pl.best[1]))
	}
	onto := unit{text: ")고양이", head: 1, name: pairName, candidates: []pairing{{shutName, 1}}}
	pl = newPlanter(unitScan(r.ix, onto), 1)
	if len(pl.transparent) != 0 {
		t.Fatal("letters after the sign", len(pl.transparent))
	}
}
