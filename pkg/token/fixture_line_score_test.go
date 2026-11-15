//ff:func feature=token type=helper control=sequence
//ff:what lineScore — what follows the tab on the printed line of a candidate (its verdict and preference as printed), from lines joined with newlines; "" when no line is that candidate
package token

import "strings"

func lineScore(out, name string) string {
	_, rest, ok := strings.Cut("\n"+out, "\n  "+name+"\t")
	if !ok {
		return ""
	}
	score, _, _ := strings.Cut(rest, "\n")
	return score
}
