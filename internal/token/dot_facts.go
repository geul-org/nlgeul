//ff:func feature=token type=reader control=iteration dimension=1
//ff:what dotFacts — the facts of the dot after the k-th run of a dotted name: the runs on its two sides, the name's last run, whether the name is numbers, and its dots
package token

import (
	"strings"

	"github.com/geul-org/koparse/pkg/rule"
)

func dotFacts(labels []string, k int) rule.Mark {
	m := rule.Mark{
		Left: sideOf(labels[k]), Right: sideOf(labels[k+1]), Ext: labels[k+1], Last: labels[len(labels)-1],
		Dots: len(labels) - 1, Run: 1, Numeric: true,
	}
	head := strings.TrimLeft(labels[0], "vV")
	m.VHead = head != labels[0] && head != ""
	for i, l := range labels {
		if i == 0 {
			l = head
		}
		m.Numeric = m.Numeric && sideOf(l)&rule.SideDigit != 0
	}
	return m
}
