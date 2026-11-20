//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what codeSpanNodes — whether the nodes of a segment follow each other from the start of its region to its end, each with its letters, spaces left out, those at its place among the letters of the sentence
package parse

import (
	"strings"
	"testing"
	"unicode"
)

func codeSpanNodes(t *testing.T, sentence string, seg Segment, letters []rune) {
	t.Helper()
	at := seg.Region.Start
	for _, n := range seg.Nodes {
		bare := strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1
			}
			return r
		}, n.Contest.Text)
		if n.Span.Start != at || bare != letterLine(n.Span, letters) {
			t.Fatal(sentence, "node", n.Span, at, n.Contest.Text)
		}
		at = n.Span.End
	}
	if len(seg.Nodes) > 0 && at != seg.Region.End {
		t.Fatal(sentence, "nodes end", at, seg.Region)
	}
}
