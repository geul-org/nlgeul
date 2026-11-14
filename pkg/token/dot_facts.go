//ff:func feature=token type=reader control=iteration dimension=1
//ff:what dotFacts — the facts of the dot after the k-th run of a dotted name: the runs on its two sides, the run to its right and the name's last run in lower case (as the lists of extensions and domains have them), and whether the name is numbers
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func dotFacts(labels []string, k int) rule.Mark {
	m := rule.Mark{
		Left: sideOf(labels[k]), Right: sideOf(labels[k+1]),
		Ext: strings.ToLower(labels[k+1]), Last: strings.ToLower(labels[len(labels)-1]),
		Run: 1, Numeric: true,
	}
	head := strings.TrimLeft(labels[0], "vV")
	for i, l := range labels {
		if i == 0 {
			l = head
		}
		m.Numeric = m.Numeric && sideOf(l)&rule.SideDigit != 0
	}
	return m
}
