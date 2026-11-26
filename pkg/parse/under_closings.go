//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what underClosings — the tails that end with the particle 는 put under a closing's last syllable as the final consonant ㄴ (비해선 = 비하+(여)서+(는), 닫을진 = 닫+을지+(는), 하다간 = 하+다가+(는)): each predicate tail that ends at j once the consonant is taken away, whose closing has letters of its own before j, judged as 는 written right after that closing is (a linking closing or one that asks: the graphs of afterClosings), with (는) as one more piece; only where no Hangul letter is written right after it (the particle ends its word: 크겐다 holds no 크+게+(는)); not for a pair the table of shortened closings gives as a closing itself (진 is 지는, 곤 is 고는), nor under the bare 어/아/여 (connNames: 먹언 is not written for 먹어는)
package parse

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The tails are read by a scan of its own whose letter before j has no
// final consonant; the stem before the tail is the unit's own.
func (sc *scan) underClosings(j int) []tail {
	if j < 2 || rule.JongOf(sc.s[j-1]) != rule.JongN || sc.right(j) != rule.RightDone {
		return nil
	}
	bare := sc.s[j-1] - rune(rule.JongN)
	ask, link := lastOfAny(askClosings, bare), lastOfAny(auxClosings, bare)
	if !ask && !link {
		return nil
	}
	sub := *sc
	sub.s = slices.Clone(sc.s)
	sub.s[j-1] = bare
	sub.placeLetters()
	e := sc.ix.named(underKey, underName)
	var out []tail
	for _, t := range sub.closings(j, false, false) {
		last := t.pieces[len(t.pieces)-1]
		pair := strings.ReplaceAll(strings.ReplaceAll(last.text, "(", ""), ")", "") + string(underKey)
		at := rule.Piece{Form: string(underKey), Before: last.names, HasStem: true, Next: -1}
		v := 0.0
		switch {
		case t.start >= j || shortOrigin(pair) || slices.Contains(connNames, last.names):
		case ask && sc.judge(sc.ix.askAux, j, j, at) > 0:
			v = 1
		case link && sc.judge(sc.ix.linkAux, j, j, at) > 0:
			v = 1
		}
		if v > 0 {
			out = append(out, joined([]tail{t}, piece{underText, underName, v, sc.ix.pref(e.count, string(underKey)), tailOf(e), false, ""})...)
		}
	}
	return out
}
