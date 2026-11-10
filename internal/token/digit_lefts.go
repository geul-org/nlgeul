//ff:func feature=token type=reader control=iteration dimension=1
//ff:what digitLefts — of the runs of pieces after a number that a piece is joined onto, those it may be joined onto when letters run on after it: only the runs that follow a run of digits, not those that follow a Hangul numeral or a Latin letter
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) digitLefts(lefts []tail, right int) []tail {
	if right != rule.RightOther {
		return lefts
	}
	var out []tail
	for _, l := range lefts {
		if l.as == "" {
			out = append(out, l)
		}
	}
	return out
}
