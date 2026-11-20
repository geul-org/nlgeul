//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what caseRows.note — what a trace adds to the reason of a case exception that applied: the word it matched and the row of that word (its line, the kind of its ground and the ground); "" for a word no row has
package parse

import "fmt"

func (c *caseRows) note(word any) string {
	for _, r := range c.Rows {
		if r.Word == word {
			return fmt.Sprintf("(%s) — 예외 표 %d: %s %s", r.Word, r.Row.Line, r.GroundKind, r.Ground)
		}
	}
	return ""
}
