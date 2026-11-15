//ff:func feature=token type=helper control=iteration dimension=2
//ff:what wayScored — wayChoices with the values of each choice after a tab, as a box prints a candidate's: the least verdict and the summed preference of the chosen candidates ("성립 %.2f 선호 %.2f")
package token

import (
	"fmt"
	"math"
)

func wayScored(seg Segment) []string {
	type choice struct {
		text          string
		verdict, pref float64
	}
	if len(seg.Nodes) == 0 {
		return nil
	}
	all := []choice{{verdict: 1}}
	for i, n := range seg.Nodes {
		var next []choice
		for _, c := range all {
			for _, b := range n.Contest.Branches {
				next = append(next, choice{c.text + map[bool]string{true: " ", false: ""}[i > 0] + b.String(), math.Min(c.verdict, b.Verdict), c.pref + b.Pref})
			}
		}
		all = next
	}
	out := make([]string, 0, len(all))
	for _, c := range all {
		out = append(out, fmt.Sprintf("%s\t성립 %.2f 선호 %.2f", c.text, c.verdict, c.pref))
	}
	return out
}
