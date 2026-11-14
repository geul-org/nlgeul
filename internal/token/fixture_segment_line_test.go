//ff:func feature=token type=helper control=iteration dimension=1
//ff:what segmentLine — a segment's nodes in short: the letters of each node with its span, joined by " " (아버지[0-3] 가방에[3-6])
package token

import (
	"fmt"
	"strings"
)

func segmentLine(sg Segment) string {
	var out []string
	for _, n := range sg.Nodes {
		out = append(out, fmt.Sprintf("%s[%d-%d]", n.Contest.Text, n.Span.Start, n.Span.End))
	}
	return strings.Join(out, " ")
}
