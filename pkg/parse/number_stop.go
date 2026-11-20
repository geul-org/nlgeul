//ff:func feature=parse type=reader control=sequence
//ff:what numberStop — the full stop right after a number that ends at j, read as a piece of the number's own word when a particle chain goes on after it (31.까지 → 31+.[마침표]+까지): where it starts and the one piece, with the verdict of the graph 마침표 that is rebutted when no number stands before the dot
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The piece costs nothing in preference, like the full stop that closes
// a word (stopAt). numEnds holds the end of a run read as a number, so a
// dot after another dot, or after a word, is no such stop.
func (sc *scan) numberStop(j int) []tail {
	if j < 2 || sc.s[j-1] != '.' {
		return nil
	}
	facts := rule.Piece{Form: ".", HasStem: true, Next: -1, Num: sc.numEnds[j-1]}
	v := sc.judge(sc.ix.numStop, j-1, j, facts)
	if v <= 0 {
		return nil
	}
	return []tail{{start: j - 1, pieces: []piece{{".", sc.ix.numStop.name, v, 0, rule.TailAux, false, ""}}}}
}
