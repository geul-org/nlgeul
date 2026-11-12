//ff:func feature=hypo type=helper control=iteration dimension=1
//ff:what Count — how many hypotheses the bundles give: the product of their numbers of readings, as a number of any size; none when there is no bundle
package hypo

import (
	"math/big"

	"github.com/geul-org/nlgeul/internal/token"
)

func Count(choices []token.Bundle) *big.Int {
	if len(choices) == 0 {
		return new(big.Int)
	}
	n := big.NewInt(1)
	for _, b := range choices {
		n.Mul(n, big.NewInt(int64(len(b.Branches))))
	}
	return n
}
