//ff:func feature=token type=reader control=iteration dimension=1
//ff:what checkCases — every graph the case exception table names is a marker graph or a rule of a sentence of the index; one that is neither is an error that names the first row of that graph, the graphs taken in the order of their names
package token

import (
	"maps"
	"slices"
)

// Asked of an index built with a word table, which has every marker
// graph (some, the 하오체 closings and the shortened ones, are made only
// with a word table) and the rules of a sentence, put last.
func (ix *index) checkCases() error {
	names := ix.markerNames()
	for _, r := range ix.sentence {
		names[r.name] = true
	}
	for _, graph := range slices.Sorted(maps.Keys(caseExceptions)) {
		if !names[graph] {
			return caseExceptions[graph][0].Row.Errf(0, "색인에 이 이름의 표지 그래프도 문장 규칙도 없다")
		}
	}
	return nil
}
