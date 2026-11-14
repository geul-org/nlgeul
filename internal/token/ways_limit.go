//ff:func feature=token type=helper control=sequence
//ff:what waysLimit — how many ways are asked of a stretch that has ways of them (a big number) when head of them are wanted (0 for all): all of them when they fit an int and none is wanted beyond head, head when there are more, and no limit (the greatest int) for a number of ways no int holds when all are wanted
package token

import (
	"math"
	"math/big"
)

func waysLimit(ways *big.Int, head int) int {
	limit := math.MaxInt
	if ways != nil && ways.IsInt64() && ways.Int64() <= int64(math.MaxInt) {
		limit = int(ways.Int64())
	}
	if head > 0 {
		limit = min(limit, head)
	}
	return limit
}
