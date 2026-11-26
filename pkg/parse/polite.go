//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what polite — the candidates of an old polite slot (오, 옵, 사오, 사옵, 자오, 자옵) right before a closing: each tail of a closing, read once more after each such slot whose letters end right before it; only with a word table (the index has no such slot without one)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The tails are those of the closings read at one place. A slot ends in
// 오 or 옵, or in 오 with a closing fused under it, so a tail after any
// other syllable is passed by.
func (sc *scan) polite(tails []tail) []tail {
	var out []tail
	for _, t := range tails {
		if t.start < 1 || rule.JungOf(sc.s[t.start-1]) != rule.JungOf(politeSoft) || rule.ChoOf(sc.s[t.start-1]) != choNone {
			continue
		}
		for _, e := range sc.ix.polite {
			out = append(out, sc.politeAt(t, e)...)
		}
	}
	return out
}
