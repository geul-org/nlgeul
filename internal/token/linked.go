//ff:func feature=token type=reader control=iteration dimension=1
//ff:what linked — the tails that begin with the link vowel 으 right before some pieces (읽+으+시, 먹+으+니): after the stem, by the syllable that stands there or, for a changed stem (들으니, 지으니), by the stem's own last syllable; and (when asked) after the slots 었, 겠 (먹었으니)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// start is where the pieces begin and after the name of the first.
func (sc *scan) linked(start int, ps []piece, after string, slots bool) []tail {
	at := start - 1
	if at < 1 || sc.s[at] != '으' {
		return nil
	}
	all := append([]piece{{"으", euName, 1, sc.ix.pref(euCount, "으"), rule.TailPred, false}}, ps...)
	facts := rule.Piece{Form: "으", HasStem: true, Prev: sc.s[at-1], Next: -1, After: after}
	if sc.judge(sc.ix.eu, at, start, facts) > 0 {
		out := []tail{{start: at, pieces: all, head: headEuForm}}
		if slots {
			out = append(out, joined(sc.slots(at, rule.SlotClosing), all...)...)
		}
		return out
	}
	var out []tail
	for _, a := range sc.ix.irrEnds[sc.s[at-1]] {
		facts.Prev = []rune(a.stem)[0]
		if sc.judge(sc.ix.eu, at, start, facts) > 0 {
			out = append(out, tail{start: at, pieces: all, head: headEuForm, irr: a.kind})
		}
	}
	return out
}
