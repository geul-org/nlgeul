//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNet — a unit of no letters has no net; a unit read by the net gives the states that end it (each at the unit's last letter), every state made once and in order, no twin; every word of every way the old reading finds is a state, and every two words side by side are a state and one of its preds
package token

import "testing"

func TestNet(t *testing.T) {
	empty, _, _ := netUnit(t, "words-stream.tsv", "가")
	empty.s = nil
	if empty.net() != nil {
		t.Fatal("no letters")
	}
	sc, _, r := netUnit(t, "words-stream.tsv", "아버지가방에들어가신다")
	lat := sc.net()
	if lat.n != 11 || len(lat.ends) == 0 || len(lat.all) != len(sc.registry().all) || lat.twins != 0 || !sc.netted {
		t.Fatal(lat.n, len(lat.ends), len(lat.all), lat.twins)
	}
	for _, e := range lat.ends {
		if e.last.end != 11 {
			t.Fatal("an end", e.last.end)
		}
	}
	old := unitScan(r.ix, unitsOf(r.ix, "아버지가방에들어가신다", false)[0])
	for _, p := range old.search() {
		if !pathInNet(lat, p) {
			t.Fatal("a way of the old reading is not in the net", p.sig)
		}
	}
}
