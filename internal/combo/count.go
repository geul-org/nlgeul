//ff:func feature=combo type=helper control=iteration dimension=1
//ff:what Count — how many combinations of candidates the contests give: the product of their numbers of candidates, as a number of any size; none when there is no contest
package combo

import (
	"math/big"

	"github.com/geul-org/nlgeul/internal/token"
)

func Count(choices []token.Contest) *big.Int {
	if len(choices) == 0 {
		return new(big.Int)
	}
	n := big.NewInt(1)
	for _, b := range choices {
		n.Mul(n, big.NewInt(int64(len(b.Branches))))
	}
	return n
}
