//ff:func feature=token type=reader control=iteration dimension=1
//ff:what checkCases — every graph the case exception table names is a marker graph of the index; one that is not is an error that names the first row of that graph, the graphs taken in the order of their names
package token

import (
	"maps"
	"slices"
)

// Asked of an index built with a word table, which has every marker
// graph: some (the 하오체 closings, the shortened ones) are made only
// with a word table.
func (ix *index) checkCases() error {
	names := ix.markerNames()
	for _, graph := range slices.Sorted(maps.Keys(caseExceptions)) {
		if !names[graph] {
			return caseExceptions[graph][0].Row.Errf(0, "색인에 이 이름의 표지 그래프가 없다")
		}
	}
	return nil
}
