//ff:func feature=parse type=reader control=sequence
//ff:what matchA — read a closing that begins with 어/아/여, as a syllable of its own or fused into the stem's last vowel (봐, 해서, 가요)
package parse

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// yo tells that 요 follows the closing. A fused closing stays in one
// piece with the stem's last syllable (해서), except under 셔, which is
// the two grammatical pieces 시 + 어.
func matchA(s []rune, end int, e ending, yo bool) (tailFact, bool) {
	start, ok := suffixAt(s, end, e.form)
	if !ok || start < 1 {
		return tailFact{}, false
	}
	if strings.ContainsRune("어아여", s[start-1]) {
		return finish(s, start-1, e, 2*(end-start)+2, "+"+string(s[start-1:end]))
	}
	if !rule.FusedVowel(s[start-1], e.form != "" || yo) {
		return tailFact{}, false
	}
	t := tailFact{rule.LeadText(e.lead) + e.form, e.kind, stemOf(s[:start]), 2*(end-start) + 1, false, e.count, affixMark(s[:start], true) + e.form}
	if s[start-1] == '셔' && start >= 2 {
		t.marked = affixMark(s[:start-1], false) + "+시+어" + e.form
	}
	return t, true
}
