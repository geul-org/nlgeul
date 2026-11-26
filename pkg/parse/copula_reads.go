//ff:func feature=parse type=reader control=sequence
//ff:what copulaReads — one predicate tail read after the copula's stem e (said or dropped) that ends at start, judged by the stem's graph, with the stem as a piece before the tail; before the closings and the slots of droppableKnown (what stands right after the stem: 네, and 시 of 누구+(이)+시+다) a dropped stem asks for a nominal of the word table before it (or a new word, before those of droppableNew: dropGate), unless the letter before it is no Hangul syllable (a run of Latin letters and digits is no row of the table), a stem dropped right after a letter that is no Hangul syllable is read only before 야, or right after a sign that closes a pair (dropBarred: '묘지'+(이)+다), and a stem dropped right after a syllable 이 with no piece fused into it asks for a nominal of the table whatever follows (아이+(이)+면: the syllable is far more often the stem itself); the graph is told what stands right after the stem (the tail's closing, or the slot before it, or the auxiliary written onto its 어/아 form: afterCopula) whether a particle is being read right after the tail (chained), and whether a number ends right before the stem (numEnds); the tail is read once more after a unit (afterUnits), after a suffix or a particle the copula follows (afterSuffixes) after a linking closing it follows (afterLinked) and after a naming closing it follows (afterNaming)
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) copulaReads(t tail, start int, e *entry, chained bool) []tail {
	p := rule.Piece{Form: "이", HasStem: start >= 1 || sc.attach, Next: rule.SlotClosing, After: sc.ix.afterCopula(t), Right: rule.RightDone}
	if chained {
		p.Right = rule.RightParticle
	}
	base := rune(0)
	if start >= 1 {
		p.Prev, p.Num = sc.s[start-1], sc.numEnds[start]
	}
	if e == sc.ix.dropped && t.base != 0 {
		base, p.Prev = t.base, t.base
	}
	if e == sc.ix.dropped && base == 0 && dropBarred(p.Prev, p.After) {
		return nil
	}
	v := sc.judge(e, start, start+len([]rune(e.form)), p)
	if v <= 0 {
		return nil
	}
	pc := piece{"이", e.name, v, sc.ix.pref(e.count, "이"), rule.TailCopula, e == sc.ix.dropped, ""}
	if pc.unseen {
		pc.text = "(이)"
	}
	read := tail{start: start, pieces: append([]piece{pc}, t.pieces...), base: base, noYo: t.noYo, close: t.close}
	if pc.unseen && slices.Contains(droppableKnown, p.After) && rule.Fits(p.Prev, rule.AfterHangul) {
		read.gate = sc.ix.dropGate(p.After)
	}
	if pc.unseen && read.gate == nil && base == 0 && p.Prev == '이' {
		read.gate = sc.ix.nominal
	}
	out := []tail{read}
	out = append(out, sc.afterUnits(read)...)
	out = append(out, sc.afterSuffixes(read)...)
	out = append(out, sc.afterLinked(read)...)
	return append(out, sc.afterNaming(read)...)
}
