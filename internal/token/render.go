//ff:func feature=token type=reader control=iteration dimension=1
//ff:what render — the lines of one contest: a single candidate on one line; more than one as a head line with what its first two branches are called, then each branch with its verdict and preference, only the first head of them when head is above 0, with a last line telling how many were left out (a count of what is shown, not of what is read)
package token

import "fmt"

func render(b Contest, head int) []string {
	if len(b.Branches) == 1 {
		return []string{b.Branches[0].String()}
	}
	out := []string{fmt.Sprintf("%s\t경합(%s / %s)", b.Text, b.Branches[0].Label(), b.Branches[1].Label())}
	for i, br := range b.Branches {
		if head > 0 && i == head {
			out = append(out, fmt.Sprintf("  (%d개 생략됨)", len(b.Branches)-head))
			break
		}
		out = append(out, fmt.Sprintf("  %s\t성립 %.2f 선호 %.2f", br.String(), br.Verdict, br.Pref))
	}
	return out
}
