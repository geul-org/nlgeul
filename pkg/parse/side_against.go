//ff:func feature=parse type=rule control=selection
//ff:what sideAgainst — the facts of its place (Mark… flags) that speak against one candidate of a quotation mark, any of which lowers it: for the opening of a mark written the same way at both ends, a letter or digit right before it or a marker that takes a quotation right after it (좋다"고, "라고); for the opening of a round mark that is also read as a closing one (swapShuts), a closing mark's place (사항‘이라고); for the closing of any quotation mark, an opening mark's place (”두, and "곰 after a space with one open); 0 for any other candidate and for a sign that is no quotation mark
package parse

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A mark written the same way at both ends is told by its place alone,
// so a word right before it is enough; a round opening mark is written
// onto the word before it often (이른바“…”), so only a closing mark's
// place speaks against its shape.
func sideAgainst(sign rune, row pairRow, name string) int {
	switch {
	case name == openName && row.shuts:
		return rule.MarkAfterWord | rule.MarkQuoted
	case name == openName && strings.ContainsRune(swapShuts, sign):
		return rule.MarkShutSide
	case name == shutName && strings.ContainsRune(quoteSigns, sign):
		return rule.MarkOpenSide
	}
	return 0
}
