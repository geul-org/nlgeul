//ff:func feature=parse type=reader control=sequence
//ff:what dropBarred — is a dropped copula stem not read right after the letter given before what stands after it: after a letter that is no Hangul syllable the stem is read as dropped before the closings of droppableRun (야) whatever the letter is, and before the other closings and slots it is dropped before only right after a sign that closes a pair (a closing bracket, a quotation mark: '묘지'+(이)+다, "…"+(이)+냐고), for what such a sign closes stands as a nominal there; after a Hangul syllable nothing is barred
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func dropBarred(prev rune, after string) bool {
	if rule.Fits(prev, rule.AfterHangul) || slices.Contains(droppableRun, after) {
		return false
	}
	return !pairRows[prev].shuts
}
