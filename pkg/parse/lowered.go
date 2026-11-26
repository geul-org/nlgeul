//ff:func feature=parse type=helper control=sequence
//ff:what lowered — a candidate whose verdict is no higher than the verdict given (that of the reading of the sign of a pair at its head), in the branch and in the score kept of its first word (a copy of the scores); the candidate itself when its verdict is already no higher
package parse

func lowered(br Branch, verdict float64) Branch {
	if br.Verdict <= verdict {
		return br
	}
	br.Verdict = verdict
	if len(br.scores) > 0 {
		br.scores = append([]wordScore{}, br.scores...)
		br.scores[0].verdict = min(br.scores[0].verdict, verdict)
	}
	return br
}
