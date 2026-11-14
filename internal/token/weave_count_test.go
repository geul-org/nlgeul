//ff:func feature=token type=helper control=sequence
//ff:what TestWeaveCount — the ways that go on from a span with the set of states come to there: one at the span that ends the portion, the ways of each span that follows with a state after one of the set, none through a span none follows; each (span, set) counted once (kept in the memo)
package token

import (
	"math/big"
	"testing"
)

func TestWeaveCount(t *testing.T) {
	w, _ := oddWeave()
	memo := map[string]*big.Int{}
	last := w.at[Span{3, 4}]
	if got := w.count(last, w.of[last], memo); got.Int64() != 1 || len(memo) != 1 {
		t.Fatal(got, len(memo))
	}
	cut := w.at[Span{2, 3}]
	if got := w.count(cut, w.of[cut], memo); got.Int64() != 1 || len(memo) != 2 {
		t.Fatal("through the span that follows", got, len(memo))
	}
	if got := w.count(cut, w.of[cut], memo); got.Int64() != 1 || len(memo) != 2 {
		t.Fatal("once", got, len(memo))
	}
	bad := w.at[Span{0, 2}]
	if got := w.count(bad, w.of[bad], memo); got.Int64() != 1 {
		t.Fatal(got)
	}
	w2, s := oddWeave()
	if got := w2.count(w2.at[Span{0, 1}], []int{0}, map[string]*big.Int{}); got.Int64() != 1 || s["c"] == nil {
		t.Fatal(got)
	}
	if got := w2.count(w2.at[Span{1, 2}], []int{}, map[string]*big.Int{}); got.Int64() != 0 {
		t.Fatal("no state, no way", got)
	}
}
