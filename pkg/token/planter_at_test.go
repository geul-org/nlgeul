//ff:func feature=token type=reader control=sequence
//ff:what TestPlanterAt — the function is worked through by the net over the sentences of the tests and made-up long units, and each is read to at least one stretch; a word of the whole unit carried over to the letters after the sign marks, at its end counted from there, that a word ending in the linking 어/아 or in a modifying closing ends there, as it does
package token

import "testing"

func TestPlanterAt(t *testing.T) {
	driveAll(t)
	sign := &path{}
	pl := &planter{head: 1, transparent: map[*path]bool{sign: true}, best: make([][]*path, 6)}
	eo := &path{start: 1, preds: []*path{sign}, last: step{end: 3, conn: true}}
	adn := &path{start: 1, preds: []*path{sign}, last: step{end: 5, adn: true}}
	pl.best[3], pl.best[5] = []*path{eo}, []*path{adn}
	ps := scanOf("가나다라마")
	if got := pl.at(3, ps); len(got) != 1 || got[0].last.end != 2 || !ps.connEnds[2] || ps.adnEnds[2] || ps.detEnds[2] {
		t.Fatal("a word that ends in the linking 어/아", got, ps.connEnds, ps.adnEnds)
	}
	if got := pl.at(5, ps); len(got) != 1 || !ps.adnEnds[4] || ps.connEnds[4] {
		t.Fatal("a word that ends in a modifying closing", got, ps.connEnds, ps.adnEnds)
	}
	if got := pl.at(4, ps); len(got) != 0 || ps.connEnds[3] {
		t.Fatal("no word ends there", got)
	}
}
