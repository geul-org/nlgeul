//ff:func feature=token type=helper control=iteration dimension=2
//ff:what contestLines — the branches of some contests (what readUnit or readOnto reads), a line each in order: the branch (Branch.String) and after a tab its verdict and preference ("성립 %.2f 선호 %.2f")
package token

import "fmt"

func contestLines(cs []Contest) []string {
	var out []string
	for _, c := range cs {
		for _, b := range c.Branches {
			out = append(out, fmt.Sprintf("%s\t성립 %.2f 선호 %.2f", b.String(), b.Verdict, b.Pref))
		}
	}
	return out
}
