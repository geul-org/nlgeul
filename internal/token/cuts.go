//ff:func feature=token type=reader control=sequence
//ff:what cuts — whether the sentence is cut at place i, and the trace lines: a sign of cutSigns with a letter right before and right after it (no space, no sentence edge), not in a span of code and not inside an address, mail address, path or option, that a graph of the sign lets stand (not a run of it, not next to another sign)
package token

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (c *cutter) cuts(i int) (bool, []string) {
	if !strings.ContainsRune(cutSigns, c.r[i]) || c.inside[i] || i == 0 || i+1 >= len(c.r) {
		return false, nil
	}
	if unicode.IsSpace(c.r[i-1]) || unicode.IsSpace(c.r[i+1]) || c.special(i) {
		return false, nil
	}
	m := rule.Mark{Left: sideOf(string(c.r[i-1])), Right: sideOf(string(c.r[i+1])), Run: 1}
	for k := i - 1; k >= 0 && c.r[k] == c.r[i]; k-- {
		m.Run++
	}
	for k := i + 1; k < len(c.r) && c.r[k] == c.r[i]; k++ {
		m.Run++
	}
	for _, r := range []rune{c.r[i-1], c.r[i+1]} {
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			m.Facts |= rule.MarkBySign
		}
	}
	var notes []string
	stands := false
	for _, e := range c.ix.marks[c.r[i]] {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, m)
		v, why := evaluate(e.g, ctx, c.trace)
		note := "통과"
		if v <= 0 {
			note = "기각: " + strings.Join(why, ", ")
		}
		if c.trace {
			notes = append(notes, fmt.Sprintf("  · %s\t%s\t%.2f\t%s", string(c.r[i]), e.name, v, note))
		}
		stands = stands || v > 0
	}
	return stands, notes
}
