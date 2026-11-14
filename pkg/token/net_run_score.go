//ff:func feature=token type=reader control=iteration dimension=1
//ff:what score — give a state its best way: the best value among the states before it that have one (a seed or the root is worth 0), the first to come on a tie, plus the best lambda of the span it stands at
package token

func (r *netRun) score(rc *reach, s *path) {
	var best float64
	var arg *path
	for _, b := range s.preds {
		v, ok := rc.of(b)
		if ok && (arg == nil || v > best) {
			best, arg = v, b
		}
	}
	sp := Span{Start: s.start, End: s.last.end}
	rc.value[s], rc.back[s] = best+r.lambdaAt(sp, rc.group[sp]), arg
}
