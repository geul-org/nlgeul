//ff:func feature=token type=reader control=iteration dimension=2
//ff:what helped — the candidates of 지 and 하 written onto a main verb's 어/아 form (지워진다, 예뻐한다): each tail that attaches to such a stem, read once more after each 어/아 form that ends right before the stem, the stem as a piece between them; only with a word table, whose predicate the main verb must be
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The tails are those of the closings read after the stem or fused
// into it. The 어/아 form before the stem is read as it is before 요:
// also when it leaves no trace in the letters (나빠진 = 나쁘+아+지+ㄴ). The main verb's own restored syllable, gate and head stay;
// what the tail asks of its stem (a predicate, a verb) is asked of 지
// or 하, not of the main verb.
func (sc *scan) helped(tails []tail) []tail {
	var out []tail
	for _, t := range tails {
		at := t.start - 1
		name, ok := helpers[sc.helper(t)]
		if !ok || !sc.ix.lexed || at < 1 || t.only != "" || t.shown != "" || t.as != "" || t.irr != 0 || (t.gate != nil && t.gate != sc.ix.pred) {
			continue
		}
		rest := append([]piece{{string(sc.helper(t)), name, 1, 0, rule.TailPred, false, ""}}, t.pieces...)
		for i := range rest {
			rest[i].tail = rule.TailPred
		}
		for _, l := range sc.closings(at, true, false) {
			if len(l.pieces) != 1 || !slices.Contains(connNames, l.pieces[0].names) {
				continue
			}
			l.pieces = append([]piece{l.pieces[0]}, rest...)
			l.noYo, l.close, l.gate = t.noYo, t.close, sc.ix.pred
			out = append(out, l)
		}
	}
	return out
}
