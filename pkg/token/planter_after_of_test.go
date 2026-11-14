//ff:func feature=token type=reader control=sequence
//ff:what TestPlanterAfterOf — a state of the whole unit that begins at the head after a state that is the sign alone is carried over to the letters after the sign (its end, and its word's end, counted from there); a state that begins elsewhere, or after any other state, is not; the units of the tests are read by the net too
package token

import "testing"

func TestPlanterAfterOf(t *testing.T) {
	driveAll(t)
	sign := &path{}
	other := &path{}
	pl := &planter{head: 2, transparent: map[*path]bool{sign: true}}
	after := &path{start: 2, preds: []*path{other, sign}, last: step{end: 5}}
	after.last.word.End = 5
	e, ok := pl.afterOf(after)
	if !ok || e.end != 3 || e.word.End != 3 {
		t.Fatal("a word right after the sign alone", e.end, ok)
	}
	if _, ok := pl.afterOf(&path{start: 2, preds: []*path{other}}); ok {
		t.Fatal("after another state")
	}
	if _, ok := pl.afterOf(&path{start: 3, preds: []*path{sign}}); ok {
		t.Fatal("a word that begins elsewhere")
	}
}
