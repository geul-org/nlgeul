//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what linked — the tails that begin with the link vowel 으 right before some pieces (읽+으+시, 먹+으+니): after the stem, by the syllable that stands there or, for a changed stem (들으니, 지으니), by the stem's own last syllable; and (when asked) after the slots 었, 겠 (먹었으니); with a word table also after a stem's ㄹ, against the tables (알으시오: 알+으+시+오), after a predicate's stem of the table and rebutted
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// start is where the pieces begin and after the name of the first.
func (sc *scan) linked(start int, ps []piece, after string, slots bool) []tail {
	at := start - 1
	if at < 1 || sc.s[at] != '으' {
		return nil
	}
	eu := piece{"으", euName, 0, sc.ix.pref(euCount, "으"), rule.TailPred, false, ""}
	facts := rule.Piece{Form: "으", HasStem: true, Prev: sc.s[at-1], Next: -1, After: after}
	if v := sc.judge(sc.ix.eu, at, start, facts); v > 0 {
		eu.verdict = v
		all := append([]piece{eu}, ps...)
		out := []tail{{start: at, pieces: all, head: headEuForm}}
		if slots {
			out = append(out, joined(sc.slots(at, rule.SlotClosing), all...)...)
		}
		return out
	}
	var out []tail
	for _, a := range sc.ix.irrEnds[sc.s[at-1]] {
		facts.Prev = []rune(a.stem)[0]
		if v := sc.judge(sc.ix.eu, at, start, facts); v > 0 {
			eu.verdict = v
			out = append(out, tail{start: at, pieces: append([]piece{eu}, ps...), head: headEuForm, irr: a.kind})
		}
	}
	facts.Prev = sc.s[at-1]
	if l := sc.ix.euLoose; l != nil && sc.judge(l, at, start, facts) > 0 {
		eu.verdict = sc.judge(l, at, start, facts)
		out = append(out, tail{start: at, pieces: append([]piece{eu}, ps...), head: headEuForm, gate: sc.ix.pred})
	}
	return out
}
