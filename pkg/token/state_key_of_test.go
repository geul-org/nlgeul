//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestStateKeyOf — the key of a candidate is its last word by everything asked of it: two candidates that differ in the words before them have the same key, and one that differs in any value of the word (verdict, preference, a flag, the cost given back, the kind) has another
package token

import "testing"

func TestStateKeyOf(t *testing.T) {
	e := wordStep(4, "w", 0.5, -2)
	e.adn, e.det, e.conn, e.back, e.kind = true, true, true, 0.7, 3
	a := grow(grow(&path{verdict: 1}, wordStep(1, "x", 1, -1)), wordStep(2, "y", 1, -1))
	b := grow(&path{verdict: 1}, wordStep(2, "z", 1, -5))
	ka, kb := stateKeyOf(grow(a, e)), stateKeyOf(grow(b, e))
	if ka != kb || ka.verdict != 0.5 || ka.pref != -2 || !ka.adn || !ka.det || !ka.conn || ka.back != 0.7 || ka.kind != 3 || ka.word != (wordKey{2, 4, "w"}) {
		t.Fatal(ka, kb)
	}
	for _, f := range []func(*step){func(s *step) { s.verdict = 0.4 }, func(s *step) { s.pref = -3 }, func(s *step) { s.adn = false }, func(s *step) { s.det = false }, func(s *step) { s.conn = false }, func(s *step) { s.back = 0 }, func(s *step) { s.kind = 0 }} {
		o := e
		f(&o)
		if stateKeyOf(grow(a, o)) == ka {
			t.Fatal("another value is another state", o)
		}
	}
}
