//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what underYo — the tails that end with 죠 as the last syllable 지 of a longer final closing with 요 fused into it (씻어야죠 = 씻+어야지+(요), 가야죠 = 가+(아)야지+(요), 가라죠 = 가+라지+(요)): each predicate tail that ends at j when 죠 is read as 지, whose closing is one of yoClosings and takes 요, with (요) as one more piece; the plain 지 is not read here (죠 is a closing of the tables written out as 지+(요))
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The tails are read by a scan of its own whose letter before j is 지;
// the stem before the tail is the unit's own.
func (sc *scan) underYo(j int) []tail {
	if j < 2 || sc.s[j-1] != yoFused {
		return nil
	}
	sub := *sc
	sub.s = slices.Clone(sc.s)
	sub.s[j-1] = yoBare
	sub.placeLetters()
	yo := piece{yoText, yoName, 1, sc.ix.pref(sc.ix.named('요', yoName).count, "요"), rule.TailPred, false, ""}
	var out []tail
	for _, t := range sub.closings(j, true, false) {
		if t.start < j && !t.noYo && slices.Contains(yoClosings, t.pieces[len(t.pieces)-1].names) {
			out = append(out, joined([]tail{t}, yo)...)
		}
	}
	return out
}
