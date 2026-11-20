//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what tildeAt — the range sign that ends at j after a number and its unit and before another number (1월~3월): the tails it closes, as the full stop closes a word, one for each candidate of the sign that stands
package parse

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) tildeAt(j int) []tail {
	if !strings.ContainsRune(tildeSigns, sc.at(j-1)) {
		return nil
	}
	lefts := sc.units(j-1, 0, false)
	if len(lefts) == 0 {
		return nil
	}
	var out []tail
	for _, r := range shownNames(sc.judgeMark([]*entry{sc.ix.tilde}, rule.Mark{Left: rule.SideHangul, Right: sideOf(string(sc.at(j))), Run: 1}, string(sc.at(j-1)))) {
		out = append(out, joined(lefts, piece{text: string(sc.at(j - 1)), names: r.name, verdict: r.v})...)
	}
	return out
}
