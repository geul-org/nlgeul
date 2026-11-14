//ff:func feature=token type=helper control=iteration dimension=1
//ff:what branchValues — the candidates of a contest that hold one word, each as its letters, verdict and preference, sorted
package token

import (
	"fmt"
	"sort"
)

func branchValues(c Contest) []string {
	var out []string
	for _, b := range c.Branches {
		if len(b.Words) == 1 {
			out = append(out, fmt.Sprintf("%s|%.6f|%.6f", b.String(), b.Verdict, b.Pref))
		}
	}
	sort.Strings(out)
	return out
}
