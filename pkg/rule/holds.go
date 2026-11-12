//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what Holds — a yes/no rule's value under its criteria: +1 when the fact holds and -1 when it does not, turned round when a Not is among the criteria
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func Holds(yes bool, specs toulmin.Specs) float64 {
	for _, s := range specs {
		if _, not := s.(*Not); not {
			return Applies(!yes)
		}
	}
	return Applies(yes)
}
