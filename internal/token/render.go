//ff:func feature=token type=reader control=iteration dimension=1
//ff:what render — the lines of one bundle: a single reading on one line; a contest as a head line with what its first two branches are called, then each branch with its verdict and preference; readings asleep as a head line with how many there are, then each
package token

import "fmt"

func render(b Bundle) []string {
	if len(b.Branches) == 1 && !b.Asleep() {
		return []string{b.Branches[0].String()}
	}
	head := fmt.Sprintf("%s\t잠듦(%d)", b.Text, len(b.Branches))
	if !b.Asleep() {
		head = fmt.Sprintf("%s\t경합(%s / %s)", b.Text, b.Branches[0].Label(), b.Branches[1].Label())
	}
	out := []string{head}
	for _, br := range b.Branches {
		out = append(out, fmt.Sprintf("  %s\t성립 %.2f 선호 %.2f", br.String(), br.Verdict, br.Pref))
	}
	return out
}
