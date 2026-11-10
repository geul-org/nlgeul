//ff:func feature=token type=rule control=sequence
//ff:what stemVerdict — the verdict of a kind's graph for a stem with what follows it, its observed count and its length; the same facts are judged once and remembered
package token

import "github.com/geul-org/koparse/pkg/rule"

func (ix *index) stemVerdict(kind, tail, count int, stem string) float64 {
	n := len([]rune(stem))
	key := [4]int{kind, tail, count, min(n, shortKey+1)}
	if v, ok := ix.seen[key]; ok {
		return v
	}
	v, _ := verdictOf(ix.kinds[kind], rule.Piece{Form: stem, Tail: tail, Count: count}, "")
	ix.seen[key] = v
	return v
}
