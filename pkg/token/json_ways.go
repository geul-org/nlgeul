//ff:func feature=token type=helper control=sequence
//ff:what jsonWays — the number of ways of a stretch for JSON: a number when it is no more than 2^53, else its digits as a string
package token

import "math/big"

func jsonWays(w *big.Int) any {
	if w == nil {
		return int64(0)
	}
	if w.IsInt64() && w.Int64() <= 1<<53 {
		return w.Int64()
	}
	return w.String()
}
