//ff:func feature=token type=helper control=sequence
//ff:what TestWaysLimit — all the ways when they fit and none is wanted beyond head; head when there are more; no limit for a number no int holds when all are wanted, and head still holds then; none known gives no limit
package token

import (
	"math"
	"math/big"
	"testing"
)

func TestWaysLimit(t *testing.T) {
	if waysLimit(big.NewInt(60), 0) != 60 || waysLimit(big.NewInt(60), 100) != 60 || waysLimit(big.NewInt(60), 3) != 3 {
		t.Fatal("fits")
	}
	huge := new(big.Int).Lsh(big.NewInt(1), 90)
	if waysLimit(huge, 0) != math.MaxInt || waysLimit(huge, 5) != 5 || waysLimit(nil, 0) != math.MaxInt {
		t.Fatal("too many for an int")
	}
}
