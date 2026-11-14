//ff:func feature=token type=helper control=selection
//ff:what regionDiff — what differs between two regions of the same stretch, "" when nothing does: the span, the nodes (span, unit, Glued, letters, and every candidate with its words, verdict, preference and number of birth, in order), the best lambda and the rest of each node, the words that follow each, how many ways there are, the places where a word came with other values; then up to deals ways are asked of both and each must be the same (rank, spans, words, lambda), and last how many walks were skipped
package token

import (
	"fmt"
	"reflect"
)

func regionDiff(a, b *region, deals int) string {
	switch {
	case a.span != b.span:
		return fmt.Sprint("span ", a.span, b.span)
	case len(a.nodes) != len(b.nodes):
		return fmt.Sprint("nodes ", len(a.nodes), len(b.nodes))
	case firstNodeDiff(a.nodes, b.nodes) != "":
		return "node " + firstNodeDiff(a.nodes, b.nodes)
	case !reflect.DeepEqual(a.best, b.best) || !reflect.DeepEqual(a.rest, b.rest) || !reflect.DeepEqual(a.next, b.next):
		return "best, rest or next"
	case a.ways.Cmp(b.ways) != 0 || a.clashes != b.clashes:
		return fmt.Sprint("ways ", a.ways, b.ways, " clashes ", a.clashes, b.clashes)
	}
	return regionRest(a, b, deals)
}
