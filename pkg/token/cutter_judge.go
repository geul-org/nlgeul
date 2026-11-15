//ff:func feature=token type=rule control=iteration dimension=1
//ff:what judge — run the graphs of a sign of a pair for one way the signs before it may stand open: the candidates that stand, each with its depth (an opening sign one deeper than what stands open, a closing sign as deep as the nearest sign of its kind that stands open, a quotation mark that closes with none of its kind open 0), the way the signs stand open after each (that of a quotation mark that closes with none of its kind open is the way as it was), and the trace lines when a trace is kept
package token

import (
	"fmt"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A closing sign closes the nearest sign of its kind and with it what
// was opened after that and not closed. A closing sign with none of its
// kind open in the way (only loneShutGraph lets one stand) closes
// nothing: its depth is counted as any closing sign's, the place of the
// nearest sign of its kind plus one, which is 0 with none found, and
// what stands open is left as it is.
func (c *cutter) judge(i int, row pairRow, m rule.Mark, way string) ([]pairing, []string, []string) {
	var candidates []pairing
	var after, notes []string
	stack := []rune(way)
	at := strings.LastIndex(way, string(row.kind))
	if at >= 0 {
		m.Facts |= rule.MarkOpened
		at = len([]rune(way[:at]))
	}
	for _, e := range c.ix.pairs[c.r[i]] {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, m)
		v, why := evaluate(e.g, ctx, c.trace)
		if c.trace {
			notes = append(notes, fmt.Sprintf("  · %s\t%s\t%.2f\t%s\t열린 기호 %q", string(c.r[i]), e.name, v, noteOf(v, why), way))
		}
		if v <= 0 {
			continue
		}
		rd, next := pairing{openName, len(stack) + 1}, way+string(row.kind)
		if e.form == shutName && at >= 0 {
			rd, next = pairing{shutName, at + 1}, string(stack[:at])
		} else if e.form == shutName {
			rd, next = pairing{shutName, at + 1}, way
		}
		candidates, after = append(candidates, rd), append(after, next)
	}
	return candidates, after, notes
}
