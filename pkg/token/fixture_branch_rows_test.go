//ff:func feature=token type=helper control=iteration dimension=1
//ff:what branchRows — the candidates of a node as one line each: the word, verdict, preference and number of birth
package token

import "fmt"

func branchRows(bs []Branch) []string {
	var out []string
	for _, b := range bs {
		out = append(out, fmt.Sprintf("%s %.4f %.4f #%d", b.String(), b.Verdict, b.Pref, bornOf(b)))
	}
	return out
}
