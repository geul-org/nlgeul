//ff:func feature=token type=reader control=iteration dimension=1
//ff:what render — the lines of one bundle: a single reading on one line; a contest as a head line with the kind of contest, then each branch with its verdict and preference
package token

import "fmt"

func render(b bundle) []string {
	if len(b.branches) == 1 {
		return []string{b.branches[0].marked}
	}
	out := []string{fmt.Sprintf("%s\t경합(%s / %s)", b.text, b.branches[0].names, b.branches[1].names)}
	for _, br := range b.branches {
		out = append(out, fmt.Sprintf("  %s\t성립 %.2f 선호 %.2f", br.marked, br.verdict, br.pref))
	}
	return out
}
