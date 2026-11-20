//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what hasMark — is one of the candidate's markers (the parts after its stem) named by one of the names given
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func hasMark(n rule.Near, names []string) bool {
	for _, m := range rule.MarksOf(n) {
		if slices.Contains(names, m) {
			return true
		}
	}
	return false
}
