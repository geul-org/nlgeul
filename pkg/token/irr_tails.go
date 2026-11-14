//ff:func feature=token type=reader control=iteration dimension=2
//ff:what irrTails — the closings that end at j and attach only after a consonant (some of them not after ㄹ), read where a vowel or ㄹ stands before them, by the own last syllable of a stem that changes there (지은 = 짓+은, 물을 = 묻+을): each stands as a tail that asks for a stem so changed
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// Such a closing takes no slot before it: the syllable it is judged by
// is the stem's.
func (sc *scan) irrTails(j int, more bool) []tail {
	var out []tail
	for _, e := range sc.entriesAt(j) {
		start, p, ok := sc.place(e, j, more)
		if !ok || e.class != classClosing || e.lead != 0 || (e.after != rule.AfterBatchim && e.after != rule.AfterHard) || start < 1 || rule.Fits(sc.s[start-1], e.after) {
			continue
		}
		for _, a := range sc.ix.irrEnds[sc.s[start-1]] {
			p.Prev = []rune(a.stem)[0]
			if v := sc.judge(e, start, j, p); v > 0 {
				pc := sc.entryPiece(e, string(sc.s[start:j]), v)
				out = append(out, tail{start: start, pieces: partsOf(pc, splits[e.name]), head: e.head, irr: a.kind, noYo: e.noYo, close: e})
			}
		}
	}
	return out
}
