//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what politeAt — the tails of one old polite slot before one closing's tail: the slot and the closing, alone, after the slots that stand before it (시, 었, 겠) and after the link vowel 으; the word before must be a predicate's stem of the word table
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The slot's letters end where the closing begins, or with the syllable
// the closing is fused into (온 = 오+ㄴ, 옵니다 = 옵+니다). The slot has
// no observed count.
func (sc *scan) politeAt(t tail, e *entry) []tail {
	letters := []rune(e.form)
	at := t.start - len(letters)
	closing, ok := sc.politeClosing(t, letters[len(letters)-1])
	if !ok || at < 1 || string(sc.s[at:t.start-1]) != string(letters[:len(letters)-1]) {
		return nil
	}
	v := sc.judge(e, at, t.start, rule.Piece{Form: e.form, HasStem: true, Prev: sc.s[at-1], Slot: e.slot, Next: rule.SlotClosing})
	if v <= 0 {
		return nil
	}
	ps := []piece{{e.form, e.name, v, sc.ix.pref(e.count, e.form), rule.TailPred, false, ""}, closing}
	out := []tail{{start: at, pieces: ps}}
	out = append(out, joined(sc.slots(at, e.slot), ps...)...)
	out = append(out, sc.linked(at, ps, e.name, true)...)
	for i := range out {
		out[i].noYo, out[i].close, out[i].gate = t.noYo, t.close, sc.ix.pred
	}
	return out
}
